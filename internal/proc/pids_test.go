package proc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPIDs(t *testing.T) {
	// no thorough test but it makes sure things are working
	pids, err := GetPIDs()
	assert.Nil(t, err)
	assert.True(t, len(pids) > 0)
}

func TestGetPIDSFromCgroup(t *testing.T) {
	// no thorough test but it makes sure things are working
	pids, err := GetPIDsFromCgroup("self")
	assert.Nil(t, err)
	assert.True(t, len(pids) > 0)
}
