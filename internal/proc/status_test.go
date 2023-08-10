// Copyright 2018 psgo authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proc

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
Seccomp_filters:        0
Cpus_allowed:   00000001
Cpus_allowed_list:      0
Mems_allowed:   1
Mems_allowed_list:      0
voluntary_ctxt_switches:        150
nonvoluntary_ctxt_switches:     545
`

func TestParseStatus(t *testing.T) {
	s, err := parseStatus("testpid", strings.Split(statusFile, "\n"))

	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "bash", s.Name)
	assert.Equal(t, "0022", s.Umask)
	assert.Equal(t, "S", s.State)
	assert.Equal(t, "17248", s.Tgid)
	assert.Equal(t, "0", s.Ngid)
	assert.Equal(t, "17200", s.PPid)
	assert.Equal(t, "0", s.TracerPid)
	assert.Equal(t, []string{"1000", "1000", "1000", "1000"}, s.Uids)
	assert.Equal(t, []string{"100", "100", "100", "100"}, s.Gids)
	assert.Equal(t, "256", s.FdSize)
	assert.Equal(t, []string{"16", "33", "100"}, s.Groups)
	assert.Equal(t, "17248", s.NStgid)
	assert.Equal(t, []string{"17248"}, s.NSpid)
	assert.Equal(t, "17248", s.NSpgid)
	assert.Equal(t, "131168", s.VMPeak)
	assert.Equal(t, "131168", s.VMSize)
	assert.Equal(t, "0", s.VMLCK)
	assert.Equal(t, "0", s.VMPin)
	assert.Equal(t, "13484", s.VMHWM)
	assert.Equal(t, "13484", s.VMRSS)
	assert.Equal(t, "10264", s.RssAnon)
	assert.Equal(t, "3220", s.RssFile)
	assert.Equal(t, "0", s.RssShmem)
	assert.Equal(t, "10332", s.VMData)
	assert.Equal(t, "136", s.VMStk)
	assert.Equal(t, "992", s.VMExe)
	assert.Equal(t, "2104", s.VMLib)
	assert.Equal(t, "76", s.VMPTE)
	assert.Equal(t, "12", s.VMPMD)
	assert.Equal(t, "0", s.VMSwap)
	assert.Equal(t, "0", s.HugetlbPages)
	assert.Equal(t, "1", s.Threads)
	assert.Equal(t, "0/3067", s.SigQ)
	assert.Equal(t, "0000000000000000", s.SigPnd)
	assert.Equal(t, "0000000000000000", s.ShdPnd)
	assert.Equal(t, "0000000000010000", s.SigBlk)
	assert.Equal(t, "0000000000384004", s.SigIgn)
	assert.Equal(t, "000000004b813efb", s.SigCgt)
	assert.Equal(t, "0000000000000000", s.CapInh)
	assert.Equal(t, "0000000000000000", s.CapPrm)
	assert.Equal(t, "0000000000000000", s.CapEff)
	assert.Equal(t, "ffffffffffffffff", s.CapBnd)
	assert.Equal(t, "0000000000000000", s.CapAmb)
	assert.Equal(t, "0", s.NoNewPrivs)
	assert.Equal(t, "0", s.Seccomp)
	assert.Equal(t, "0", s.SeccompFilters)
	assert.Equal(t, "00000001", s.CpusAllowed)
	assert.Equal(t, "0", s.CpusAllowedList)
	assert.Equal(t, "1", s.MemsAllowed)
	assert.Equal(t, "0", s.MemsAllowedList)
	assert.Equal(t, "150", s.VoluntaryCtxtSwitches)
	assert.Equal(t, "545", s.NonvoluntaryCtxtSwitches)
}
