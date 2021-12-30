package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNumbers(t *testing.T) {
	exp := []int{3,4,3,1,2}
	numbers := ParseNumbers("3,4,3,1,2")
	assert.ElementsMatch(t, exp, numbers)
}

func TestLanternfishCycle(t *testing.T) {
	state := []int{3,4,3,1,2}

	state = LanternfishCycle(state)
	assert.ElementsMatch(t, []int{2,3,2,0,1}, state)

	state = LanternfishCycle(state)
	assert.ElementsMatch(t, []int{1,2,1,6,0,8}, state)

	state = LanternfishCycle(state)
	assert.ElementsMatch(t, []int{0,1,0,5,6,7,8}, state)

	state = LanternfishCycle(state)
	assert.ElementsMatch(t, []int{6,0,6,4,5,6,7,8,8}, state)
}

func TestLanternfishCycleN(t *testing.T) {
	start := []int{3,4,3,1,2}

	out18 := LanternfishCycleN(start, 18)
	assert.Equal(t, 26, len(out18))

	out80 := LanternfishCycleN(start, 80)
	assert.Equal(t, 5934, len(out80))
}