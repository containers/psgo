package dev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTTYs(t *testing.T) {
	// no thorough test but it makes sure things are working
	devs, err := getTTYs()
	assert.Nil(t, err)
	assert.NotNil(t, devs)
}
