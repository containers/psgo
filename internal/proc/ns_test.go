package proc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePIDNamespace(t *testing.T) {
	// no thorough test but it makes sure things are working
	pidNS, err := ParsePIDNamespace("self")
	assert.Nil(t, err)
	assert.True(t, len(pidNS) > 0)
}
