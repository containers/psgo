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

package host

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClockTicks(t *testing.T) {
	// no thorough test but it makes sure things are working
	ticks, err := ClockTicks()
	assert.Nil(t, err)
	assert.True(t, ticks > 0)
}

func TestBootTime(t *testing.T) {
	// no thorough test but it makes sure things are working
	btime, err := BootTime()
	assert.Nil(t, err)
	assert.True(t, btime > 0)
}
