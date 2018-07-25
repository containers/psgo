package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClockTicks(t *testing.T) {
	// no thorough test but it makes sure things are working
	ticks := ClockTicks()
	assert.True(t, ticks > 0)
}

func TestBootTime(t *testing.T) {
	// no thorough test but it makes sure things are working
	btime, err := BootTime()
	assert.Nil(t, err)
	assert.True(t, btime > 0)
}
