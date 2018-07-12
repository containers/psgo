package ps

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/unix"
)

// processFunc is used to map a given aixFormatDescriptor to a corresponding
// function extracting the desired data from a process.Process.
type processFunc func(*process.Process) (string, error)

// aixFormatDescriptor as mentioned in the ps(1) manpage.  A given descriptor
// can either be specified via its code (e.g., "%C") or its normal representation
// (e.g., "pcpu") and will be printed under its corresponding header (e.g, "%CPU").
type aixFormatDescriptor struct {
	code   string
	normal string
	header string
	procFn processFunc
}

// DefaultFormat is the `ps -ef` compatible default format.
const DefaultFormat = "user,pid,ppid,pcpu,etime,tty,time,comm"

// processDescriptors calls each `procFn` of all formatDescriptors on each
// process and returns an array of tab-separated strings.
func processDescriptors(formatDescriptors []aixFormatDescriptor, processes []*process.Process) ([]string, error) {
	data := []string{}
	// create header
	headerArr := []string{}
	for _, desc := range formatDescriptors {
		headerArr = append(headerArr, desc.header)
	}
	data = append(data, strings.Join(headerArr, "\t"))

	// dispatch all descriptor functions on each process
	for _, proc := range processes {
		pData := []string{}
		for _, desc := range formatDescriptors {
			dataStr, err := desc.procFn(proc)
			if err != nil {
				return nil, err
			}
			pData = append(pData, dataStr)
		}
		data = append(data, strings.Join(pData, "\t"))
	}

	return data, nil
}

// JoinNamespaceAndProcessInfo has the same semantics as ProcessInfo but joins
// the mount namespace of the specified pid before extracting data from `/proc`.
func JoinNamespaceAndProcessInfo(pid, format string) ([]string, error) {
	var (
		data    []string
		dataErr error
		wg      sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.LockOSThread()

		fd, err := os.Open(fmt.Sprintf("/proc/%s/ns/mnt", pid))
		if err != nil {
			dataErr = err
			return
		}
		defer fd.Close()

		// create a new mountns on the current thread
		if err = unix.Unshare(unix.CLONE_NEWNS); err != nil {
			dataErr = err
			return
		}
		unix.Setns(int(fd.Fd()), unix.CLONE_NEWNS)
		data, dataErr = ProcessInfo(format)
	}()
	wg.Wait()

	return data, dataErr
}

// ProcessInfo returns the process information of all processes in the current
// mount namespace. The input format must be a comma-separated list of
// supported AIX format descriptors.  If the input string is empty, the
// DefaultFormat is used.
// The return value is an array of tab-separated strings, to easily use the
// output for column-based formatting (e.g., with the `text/tabwriter` package).
func ProcessInfo(format string) ([]string, error) {
	if len(format) == 0 {
		format = DefaultFormat
	}

	formatDescriptors, err := parseDescriptors(format)
	if err != nil {
		return nil, err
	}

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	return processDescriptors(formatDescriptors, processes)
}

// parseDescriptors parses the input string and returns a correspodning array
// of aixFormatDescriptors, which are expected to be separated by commas.
// The input format is "desc1, desc2, ..., desN" where a given descriptor can be
// specified both, in the code and in the normal form.  A concrete example is
// "pid, %C, nice, %a".
func parseDescriptors(input string) ([]aixFormatDescriptor, error) {
	formatDescriptors := []aixFormatDescriptor{}
	for _, s := range strings.Split(input, ",") {
		s = strings.TrimSpace(s)
		found := false
		for _, d := range descriptors {
			if s == d.code || s == d.normal {
				formatDescriptors = append(formatDescriptors, d)
				found = true
			}
		}
		if !found {
			return nil, errors.Wrapf(ErrUnkownDescriptor, "'%s'", s)
		}
	}
	return formatDescriptors, nil
}

// processPCPU returns how many percent of the CPU time process p uses as
// a three digit float as string.
func processPCPU(p *process.Process) (string, error) {
	f64, err := p.CPUPercent()
	if err != nil {
		return "", err
	}
	pcpu := strconv.FormatFloat(f64, 'f', 3, 64)
	return pcpu, nil
}

// processGROUP returns the effective group ID of the process.  This will be
// the textual group ID, if it can be optained, or a decimal representation
// otherwise.
func processGROUP(p *process.Process) (string, error) {
	gids, err := p.Gids()
	if err != nil {
		return "", err
	}
	if len(gids) > 1 {
		gid := strconv.Itoa(int(gids[1]))
		g, err := user.LookupGroupId(gid)
		if err != nil {
			switch err.(type) {
			case user.UnknownGroupError:
				return gid, nil
			default:
				return "", err
			}
		}
		return g.Name, nil
	}
	return "", nil
}

// processPPID returns the parent process ID of process p.
func processPPID(p *process.Process) (string, error) {
	ppid, err := p.Ppid()
	if err != nil {
		return "", nil
	}
	return strconv.FormatInt(int64(ppid), 10), nil
}

// processUSER returns the effective user name of the process.  This will be
// the textual group ID, if it can be optained, or a decimal representation
// otherwise.
func processUSER(p *process.Process) (string, error) {
	uids, err := p.Uids()
	if err != nil {
		return "", err
	}
	if len(uids) > 1 {
		uid := strconv.Itoa(int(uids[1]))
		if uid == "0" {
			return "root", nil
		}
		u, err := user.LookupId(uid)
		if err != nil {
			switch err.(type) {
			case user.UnknownUserError:
				return uid, nil
			default:
				return "", err
			}
		}
		return u.Username, nil
	}
	return "", nil
}

// processName returns the name of process p in the format "[$name]".
func processName(p *process.Process) (string, error) {
	name, err := p.Name()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("[%s]", name), nil
}

// processARGS returns the command of p with all its arguments.
func processARGS(p *process.Process) (string, error) {
	args, err := p.CmdlineSlice()
	if err != nil {
		return "", nil
	}
	// ps (1) returns "[$name]" if command/args are empty
	if len(args) == 0 {
		return processName(p)
	}
	return strings.Join(args, " "), nil
}

// processCOMM returns the command name (i.e., executable name) of process p.
func processCOMM(p *process.Process) (string, error) {
	args, err := p.CmdlineSlice()
	if err != nil {
		return "", nil
	}
	// ps (1) returns "[$name]" if command/args are empty
	if len(args) == 0 {
		return processName(p)
	}
	spl := strings.Split(args[0], "/")
	return spl[len(spl)-1], nil
}

// processRGROUP returns the real group ID of the process.  This will be
// the textual group ID, if it can be optained, or a decimal representation
// otherwise.
func processRGROUP(p *process.Process) (string, error) {
	gids, err := p.Gids()
	if err != nil {
		return "", err
	}
	if len(gids) > 0 {
		gid := strconv.Itoa(int(gids[0]))
		g, err := user.LookupGroupId(gid)
		if err != nil {
			switch err.(type) {
			case user.UnknownGroupError:
				return gid, nil
			default:
				return "", err
			}
		}
		return g.Name, nil
	}
	return "", nil
}

// processNICE returns the nice value of process p.
func processNICE(p *process.Process) (string, error) {
	nice, err := p.Nice()
	if err != nil {
		return "", nil
	}
	return strconv.Itoa(int(nice)), nil
}

// processPID returns the process ID of process p.
func processPID(p *process.Process) (string, error) {
	return strconv.Itoa(int(p.Pid)), nil
}

// processPGID returns the process group ID of process p.
//
// TODO: currently, that's not supported by github.com/shirou/gopsutil/process
// so we have to extract the data ourselves.  That's bad, because it may look
// very different on darwin.  Task: get that upstream.
func processPGID(p *process.Process) (string, error) {
	path := fmt.Sprintf("/proc/%d/stat", p.Pid)
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error extracting process group ID: %v", err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	spl := strings.Split(scanner.Text(), " ")
	if len(spl) < 37 {
		return "", fmt.Errorf("unexpected data from '%s': %v", path, spl)
	}
	// the 5th field is the process group ID
	return spl[4], nil
}

// processETIME returns the elapsed time since the process was started, in the
// form [[DD-]hh:]mm:ss.
//
// TODO: golang's time.Duration doesn't support the upper format, so we have to
// implement on our on (if desired).  The current format looks like
// "9h49m21.457862147s" which is human readable but not what users of ps (1) are
// accustomed to.
func processETIME(p *process.Process) (string, error) {
	// created time of the process in milliseconds since the epoch
	cTime, err := p.CreateTime()
	if err != nil {
		return "", nil
	}
	created := time.Unix(0, cTime*int64(time.Millisecond))
	now := time.Now()
	elapsed := now.Sub(created)
	return fmt.Sprintf("%v", elapsed), nil
}

// processRUSER returns the effective user name of the process.  This will be
// the textual group ID, if it can be optained, or a decimal representation
// otherwise.
func processRUSER(p *process.Process) (string, error) {
	uids, err := p.Uids()
	if err != nil {
		return "", err
	}
	if len(uids) > 0 {
		uid := strconv.Itoa(int(uids[0]))
		if uid == "0" {
			return "root", nil
		}
		u, err := user.LookupId(uid)
		if err != nil {
			switch err.(type) {
			case user.UnknownUserError:
				return uid, nil
			default:
				return "", err
			}
		}
		return u.Username, nil
	}
	return "", nil
}

// processTIME returns the cumulative CPU time of process p, in the form
// "[DD-]HH:MM:SS".
//
// TODO: golang's time.Duration doesn't support the upper format, so we have to
// implement on our on (if desired).  The current format looks like
// "9h49m21.457862147s" which is human readable but not what users of ps (1) are
func processTIME(p *process.Process) (string, error) {
	tStat, err := p.Times()
	if err != nil {
		return "", nil
	}

	total := tStat.Total() // float64 in seconds
	dur, err := time.ParseDuration(fmt.Sprintf("%fs", total))
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%v", dur), nil
}

// processTTY returns the controlling tty (terminal) of process p.
func processTTY(p *process.Process) (string, error) {
	tty, err := p.Terminal()
	if err != nil {
		return "", err
	}
	if len(tty) == 0 {
		tty = "?"
	}
	return strings.TrimPrefix(tty, "/"), nil
}

// processVSZ returns the virtual memory size of process p in KiB (1024-byte
// units).
func processVSZ(p *process.Process) (string, error) {
	mStat, err := p.MemoryInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", mStat.VMS/1024), nil
}

var (
	// ErrUnkownDescriptor is returned when an unknown descriptor is parsed.
	ErrUnkownDescriptor = errors.New("unkown descriptor")

	descriptors = []aixFormatDescriptor{
		{
			code:   "%C",
			normal: "pcpu",
			header: "%CPU",
			procFn: processPCPU,
		},
		{
			code:   "%G",
			normal: "group",
			header: "GROUP",
			procFn: processGROUP,
		},
		{
			code:   "%P",
			normal: "ppid",
			header: "PPID",
			procFn: processPPID,
		},
		{
			code:   "%U",
			normal: "user",
			header: "USER",
			procFn: processUSER,
		},
		{
			code:   "%a",
			normal: "args",
			header: "COMMAND",
			procFn: processARGS,
		},
		{
			code:   "%c",
			normal: "comm",
			header: "COMMAND",
			procFn: processCOMM,
		},
		{
			code:   "%g",
			normal: "rgroup",
			header: "RGROUP",
			procFn: processRGROUP,
		},
		{
			code:   "%n",
			normal: "nice",
			header: "NI",
			procFn: processNICE,
		},
		{
			code:   "%p",
			normal: "pid",
			header: "PID",
			procFn: processPID,
		},
		{
			code:   "%r",
			normal: "pgid",
			header: "PGID",
			procFn: processPGID,
		},
		{
			code:   "%t",
			normal: "etime",
			header: "ELAPSED",
			procFn: processETIME,
		},
		{
			code:   "%u",
			normal: "ruser",
			header: "RUSER",
			procFn: processRUSER,
		},
		{
			code:   "%x",
			normal: "time",
			header: "TIME",
			procFn: processTIME,
		},
		{
			code:   "%y",
			normal: "tty",
			header: "TTY",
			procFn: processTTY,
		},
		{
			code:   "%z",
			normal: "vsz",
			header: "VSZ",
			procFn: processVSZ,
		}}
)
