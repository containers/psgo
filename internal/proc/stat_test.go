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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	statFile      = "31404 (gedit) R 2109 2128 2128 0 -1 4194304 13153 328 0 0 590 55 0 0 20 0 6 0 1331588 419667968 19515 18446744073709551615 94120519110656 94120519115256 140737253236304 0 0 0 0 4096 0 0 0 0 17 2 0 0 62588346 0 0 94120521215368 94120521216168 94120544436224 140737253242331 140737253242369 140737253242369 140737253244905 0"
	statFileSpace = "31405 (ge d it) R 2109 2128 2128 0 -1 4194304 13153 328 0 0 590 55 0 0 20 0 6 0 1331588 419667968 19515 18446744073709551615 94120519110656 94120519115256 140737253236304 0 0 0 0 4096 0 0 0 0 17 2 0 0 62588346 0 0 94120521215368 94120521216168 94120544436224 140737253242331 140737253242369 140737253242369 140737253244905 0"
	statFileParen = "31406 (ged)it) R 2109 2128 2128 0 -1 4194304 13153 328 0 0 590 55 0 0 20 0 6 0 1331588 419667968 19515 18446744073709551615 94120519110656 94120519115256 140737253236304 0 0 0 0 4096 0 0 0 0 17 2 0 0 62588346 0 0 94120521215368 94120521216168 94120544436224 140737253242331 140737253242369 140737253242369 140737253244905 0"
)

func testReadStat(file string) (string, error) {
	switch file {
	case "/proc/31404/stat":
		return statFile, nil
	case "/proc/31405/stat":
		return statFileSpace, nil
	case "/proc/31406/stat":
		return statFileParen, nil
	}
	return "", errors.New("unimplemented test case")
}

func TestParseStat(t *testing.T) {
	readStat = testReadStat

	s, err := ParseStat("31404")

	assert.Nil(t, err)
	assert.NotNil(t, s)

	assert.Equal(t, "31404", s.Pid)
	assert.Equal(t, "gedit", s.Comm)
	assert.Equal(t, "R", s.State)
	assert.Equal(t, "2109", s.Ppid)
	assert.Equal(t, "2128", s.Pgrp)
	assert.Equal(t, "2128", s.Session)
	assert.Equal(t, "0", s.TtyNr)
	assert.Equal(t, "-1", s.Tpgid)
	assert.Equal(t, "4194304", s.Flags)
	assert.Equal(t, "13153", s.Minflt)
	assert.Equal(t, "328", s.Cminflt)
	assert.Equal(t, "0", s.Majflt)
	assert.Equal(t, "0", s.Cmajflt)
	assert.Equal(t, "590", s.Utime)
	assert.Equal(t, "55", s.Stime)
	assert.Equal(t, "0", s.Cutime)
	assert.Equal(t, "0", s.Cstime)
	assert.Equal(t, "20", s.Priority)
	assert.Equal(t, "0", s.Nice)
	assert.Equal(t, "6", s.NumThreads)
	assert.Equal(t, "0", s.Itrealvalue)
	assert.Equal(t, "1331588", s.Starttime)
	assert.Equal(t, "419667968", s.Vsize)

	s2, err := ParseStat("31405")

	assert.Nil(t, err)
	assert.NotNil(t, s2)

	assert.Equal(t, "ge d it", s2.Comm)

	s3, err := ParseStat("31406")

	assert.Nil(t, err)
	assert.NotNil(t, s3)

	assert.Equal(t, "ged)it", s3.Comm)
}
