package proc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAttrCurrent(t *testing.T) {
	// no thorough test but it makes sure things are working
	_, err := ParseAttrCurrent("self")
	assert.Nil(t, err)
}
