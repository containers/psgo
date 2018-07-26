package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	// no thorough test but it makes sure things are working
	p, err := New("self")
	assert.Nil(t, err)

	assert.NotNil(t, p.Stat)
	assert.NotNil(t, p.Status)
	assert.NotNil(t, p.CmdLine)
	assert.True(t, len(p.PidNS) > 0)
	assert.True(t, len(p.Label) > 0)

	err = p.SetHostData()
	assert.Nil(t, err)
	assert.True(t, len(p.Huser) > 0)
	assert.True(t, len(p.Hgroup) > 0)
}
