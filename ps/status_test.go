package ps

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// copied from proc(5) manpage
var statusFile = `
Name:   bash
Umask:  0022
State:  S (sleeping)
Tgid:   17248
Ngid:   0
Pid:    17248
PPid:   17200
TracerPid:      0
Uid:    1000    1000    1000    1000
Gid:    100     100     100     100
FDSize: 256
Groups: 16 33 100
NStgid: 17248
NSpid:  17248
NSpgid: 17248
NSsid:  17200
VmPeak:     131168 kB
VmSize:     131168 kB
VmLck:           0 kB
VmPin:           0 kB
VmHWM:       13484 kB
VmRSS:       13484 kB
RssAnon:     10264 kB
RssFile:      3220 kB
RssShmem:        0 kB
VmData:      10332 kB
VmStk:         136 kB
VmExe:         992 kB
VmLib:        2104 kB
VmPTE:          76 kB
VmPMD:          12 kB
VmSwap:          0 kB
HugetlbPages:          0 kB        # 4.4
Threads:        1
SigQ:   0/3067
SigPnd: 0000000000000000
ShdPnd: 0000000000000000
SigBlk: 0000000000010000
SigIgn: 0000000000384004
SigCgt: 000000004b813efb
CapInh: 0000000000000000
CapPrm: 0000000000000000
CapEff: 0000000000000000
CapBnd: ffffffffffffffff
CapAmb:   0000000000000000
NoNewPrivs:     0
Seccomp:        0
Cpus_allowed:   00000001
Cpus_allowed_list:      0
Mems_allowed:   1
Mems_allowed_list:      0
voluntary_ctxt_switches:        150
nonvoluntary_ctxt_switches:     545
`

func testReadStatus(_ string) ([]string, error) {
	return strings.Split(statusFile, "\n"), nil
}

func TestParseStatus(t *testing.T) {
	readStatus = testReadStatus

	s, err := parseStatus("")

	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "bash", s.name)
	assert.Equal(t, "0022", s.umask)
	assert.Equal(t, "S", s.state)
	assert.Equal(t, "17248", s.tgid)
	assert.Equal(t, "0", s.ngid)
	assert.Equal(t, "17200", s.pPid)
	assert.Equal(t, "0", s.tracerPid)
	assert.Equal(t, []string{"1000", "1000", "1000", "1000"}, s.uids)
	assert.Equal(t, []string{"100", "100", "100", "100"}, s.gids)
	assert.Equal(t, "256", s.fdSize)
	assert.Equal(t, []string{"16", "33", "100"}, s.groups)
	assert.Equal(t, "17248", s.nStgid)
	assert.Equal(t, "17248", s.nSpid)
	assert.Equal(t, "17248", s.nSpgid)
	assert.Equal(t, "131168", s.vmPeak)
	assert.Equal(t, "131168", s.vmSize)
	assert.Equal(t, "0", s.vmLCK)
	assert.Equal(t, "0", s.vmPin)
	assert.Equal(t, "13484", s.vmHWM)
	assert.Equal(t, "13484", s.vmRSS)
	assert.Equal(t, "10264", s.rssAnon)
	assert.Equal(t, "3220", s.rssFile)
	assert.Equal(t, "0", s.rssShmem)
	assert.Equal(t, "10332", s.vmData)
	assert.Equal(t, "136", s.vmStk)
	assert.Equal(t, "992", s.vmExe)
	assert.Equal(t, "2104", s.vmLib)
	assert.Equal(t, "76", s.vmPTE)
	assert.Equal(t, "12", s.vmPMD)
	assert.Equal(t, "0", s.vmSwap)
	assert.Equal(t, "0", s.hugetlbPages)
	assert.Equal(t, "1", s.threads)
	assert.Equal(t, "0/3067", s.sigQ)
	assert.Equal(t, "0000000000000000", s.sigPnd)
	assert.Equal(t, "0000000000000000", s.shdPnd)
	assert.Equal(t, "0000000000010000", s.sigBlk)
	assert.Equal(t, "0000000000384004", s.sigIgn)
	assert.Equal(t, "000000004b813efb", s.sigCgt)
	assert.Equal(t, "0000000000000000", s.capInh)
	assert.Equal(t, "0000000000000000", s.capPrm)
	assert.Equal(t, "0000000000000000", s.capEff)
	assert.Equal(t, "ffffffffffffffff", s.capBnd)
	assert.Equal(t, "0000000000000000", s.capAmb)
	assert.Equal(t, "0", s.noNewPrivs)
	assert.Equal(t, "0", s.seccomp)
	assert.Equal(t, "00000001", s.cpusAllowed)
	assert.Equal(t, "0", s.cpusAllowedList)
	assert.Equal(t, "1", s.memsAllowed)
	assert.Equal(t, "0", s.memsAllowedList)
	assert.Equal(t, "150", s.voluntaryCtxtSwitches)
	assert.Equal(t, "545", s.nonvoluntaryCtxtSwitches)
}
