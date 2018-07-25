package proc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCmdLine(t *testing.T) {
	// no thorough test but it makes sure things are working
	_, err := ParseCmdLine("self")
	assert.Nil(t, err)
}
