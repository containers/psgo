package ps

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var statFile = "31404 (gedit) R 2109 2128 2128 0 -1 4194304 13153 328 0 0 590 55 0 0 20 0 6 0 1331588 419667968 19515 18446744073709551615 94120519110656 94120519115256 140737253236304 0 0 0 0 4096 0 0 0 0 17 2 0 0 62588346 0 0 94120521215368 94120521216168 94120544436224 140737253242331 140737253242369 140737253242369 140737253244905 0"

func testReadStat(_ string) ([]string, error) {
	return strings.Fields(statFile), nil
}

func TestParseStat(t *testing.T) {
	readStat = testReadStat

	s, err := parseStat("")

	assert.Nil(t, err)
	assert.NotNil(t, s)

	assert.Equal(t, "31404", s.pid)
	assert.Equal(t, "(gedit)", s.comm)
	assert.Equal(t, "R", s.state)
	assert.Equal(t, "2109", s.ppid)
	assert.Equal(t, "2128", s.pgrp)
	assert.Equal(t, "2128", s.session)
	assert.Equal(t, "0", s.ttyNr)
	assert.Equal(t, "-1", s.tpgid)
	assert.Equal(t, "4194304", s.flags)
	assert.Equal(t, "13153", s.minflt)
	assert.Equal(t, "328", s.cminflt)
	assert.Equal(t, "0", s.majflt)
	assert.Equal(t, "0", s.cmajflt)
	assert.Equal(t, "590", s.utime)
	assert.Equal(t, "55", s.stime)
	assert.Equal(t, "0", s.cutime)
	assert.Equal(t, "0", s.cstime)
	assert.Equal(t, "20", s.priority)
	assert.Equal(t, "0", s.nice)
	assert.Equal(t, "6", s.numThreads)
	assert.Equal(t, "0", s.itrealvalue)
	assert.Equal(t, "1331588", s.starttime)
	assert.Equal(t, "419667968", s.vsize)
}
