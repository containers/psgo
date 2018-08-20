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

	comm, err := processARGS(&p)
	assert.Nil(t, err)
	assert.Equal(t, "[foo-bar]", comm)

	p = process.Process{
		CmdLine: []string{"/usr/bin/foo-bar -flag1 -flag2"},
	}

	comm, err = processARGS(&p)
	assert.Nil(t, err)
	assert.Equal(t, "/usr/bin/foo-bar -flag1 -flag2", comm)
}

func TestProcessCOMM(t *testing.T) {
	p := process.Process{
		Status: proc.Status{
			Name: "foo-bar",
		},
		CmdLine: []string{""},
	}

	comm, err := processCOMM(&p)
	assert.Nil(t, err)
	assert.Equal(t, "[foo-bar]", comm)

	p = process.Process{
		CmdLine: []string{"/usr/bin/foo-bar"},
	}

	comm, err = processCOMM(&p)
	assert.Nil(t, err)
	assert.Equal(t, "foo-bar", comm)
}
