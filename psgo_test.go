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

package psgo

import (
	"testing"

	"github.com/containers/psgo/internal/proc"
	"github.com/containers/psgo/internal/process"
	"github.com/stretchr/testify/assert"
)

func TestProcessARGS(t *testing.T) {
	p := process.Process{
		Status: proc.Status{
			Name: "foo-bar",
		},
		CmdLine: []string{""},
	}

	ctx := new(psContext)
	comm, err := processARGS(&p, ctx)
	assert.Nil(t, err)
	assert.Equal(t, "[foo-bar]", comm)

	p = process.Process{
		CmdLine: []string{"/usr/bin/foo-bar -flag1 -flag2"},
	}

	comm, err = processARGS(&p, ctx)
	assert.Nil(t, err)
	assert.Equal(t, "/usr/bin/foo-bar -flag1 -flag2", comm)
}

func TestProcessCOMM(t *testing.T) {
	p := process.Process{
		Stat: proc.Stat{
			Comm: "foo-bar",
		},
		CmdLine: []string{""},
	}

	ctx := new(psContext)
	comm, err := processCOMM(&p, ctx)
	assert.Nil(t, err)
	assert.Equal(t, "foo-bar", comm)

	p = process.Process{
		Stat: proc.Stat{
			Comm: "foo-bar",
		},
		CmdLine: []string{"/usr/bin/foo-bar"},
	}

	comm, err = processCOMM(&p, ctx)
	assert.Nil(t, err)
	assert.Equal(t, "foo-bar", comm)
}
