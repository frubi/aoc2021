package main

import (
	"github.com/stretchr/testify/assert"
	"testing"	
)

func TestReadNumbers(t *testing.T) {
	expect := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	
	nums, err := ReadNumbers("day01_test.txt")
	assert.NoError(t, err)
	assert.ElementsMatch(t, expect, nums)
}

func TestDepthIncreases(t *testing.T) {
	nums, err := ReadNumbers("day01_test.txt")
	assert.NoError(t, err)

	o := DepthIncreases(nums)
	assert.Equal(t, 7, o)
}

func TestDepthIncreasesWindowed(t *testing.T) {
	nums, err := ReadNumbers("day01_test.txt")
	assert.NoError(t, err)

	o := DepthIncreasesWindowed(nums)
	assert.Equal(t, 5, o)
}


