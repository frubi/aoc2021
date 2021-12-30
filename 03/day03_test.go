package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMostCommonBit(t *testing.T) {
	lines, err := ReadLines("day03_test.txt")
	assert.NoError(t, err)

	c := MostCommonBit(lines, 0)
	assert.Equal(t, 1, c)

	c = MostCommonBit(lines, 1)
	assert.Equal(t, 0, c)
}

func TestLeastCommonBit(t *testing.T) {
	lines, err := ReadLines("day03_test.txt")
	assert.NoError(t, err)

	c := LeastCommonBit(lines, 0)
	assert.Equal(t, 0, c)

	c = LeastCommonBit(lines, 1)
	assert.Equal(t, 1, c)
}

func TestBuildReport(t *testing.T) {
	lines, err := ReadLines("day03_test.txt")
	assert.NoError(t, err)

	gamma := BuildReport(lines, MostCommonBit)
	fmt.Printf("gamma = %b %d\n", gamma, gamma)
	assert.Equal(t, 22, gamma)

	epsilon := BuildReport(lines, LeastCommonBit)
	fmt.Printf("epsilon = %b %d\n", epsilon, epsilon)
	assert.Equal(t, 9, epsilon)
}

func TestFilterReport(t *testing.T) {
	lines, err := ReadLines("day03_test.txt")
	assert.NoError(t, err)

	oxygenStr := FilterReport(lines, MostCommonBit, 0)
	oxygen := SafeBinToDec(oxygenStr)
	fmt.Printf("oxygen = %d (from %q)\n", oxygen, oxygenStr)
	assert.Equal(t, 23, oxygen)

	co2Str := FilterReport(lines, LeastCommonBit, 0)
	co2 := SafeBinToDec(co2Str)
	fmt.Printf("co2 = %d (from %q)\n", co2, co2Str)
	assert.Equal(t, 10, co2)

}