package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePoint(t *testing.T) {
	point := ParsePoint("9,4")
	assert.Equal(t, Point{X: 9, Y: 4}, point)
}

func TestParseLine(t *testing.T) {
	line := ParseLine("9,4 -> 3,4")
	assert.Equal(t, Line{Start: Point{X: 9, Y: 4}, End: Point{X: 3, Y: 4}}, line)
}

func TestGetLinePointsX(t *testing.T) {
	// An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
	line := ParseLine("1,1 -> 1,3")

	exp := []Point{{1, 1}, {1, 2}, {1, 3}}
	points := GetLinePoints(line, true)
	assert.ElementsMatch(t, exp, points)
}

func TestGetLinePointsY(t *testing.T) {
	// An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
	line := ParseLine("9,7 -> 7,7")

	exp := []Point{{9, 7}, {8, 7}, {7, 7}}
	points := GetLinePoints(line, true)
	assert.ElementsMatch(t, exp, points)
}

func TestGetLinePointsDown(t *testing.T) {
	// An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
	line := ParseLine("1,1 -> 3,3")

	exp := []Point{{1, 1}, {2, 2}, {3, 3}}
	points := GetLinePoints(line, true)
	assert.ElementsMatch(t, exp, points)
}

func TestGetLinePointsUp(t *testing.T) {
	// An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
	line := ParseLine("9,7 -> 7,9")

	exp := []Point{{9, 7}, {8, 8}, {7, 9}}
	points := GetLinePoints(line, true)
	assert.ElementsMatch(t, exp, points)
}

func TestGetBounds(t *testing.T) {
	strs, err := ReadLines("day05_test.txt")
	assert.NoError(t, err)

	lines := ParseLines(strs)
	min, max := GetBounds(lines)

	assert.Equal(t, 0, min.X)
	assert.Equal(t, 0, min.Y)

	assert.Equal(t, 9, max.X)
	assert.Equal(t, 9, max.Y)
}

func TestGetOverlaps(t *testing.T) {
	strs, err := ReadLines("day05_test.txt")
	assert.NoError(t, err)

	lines := ParseLines(strs)
	overlaps := GetOverlaps(lines, false)

	// Punkte mit nur einer Linie
	singlePoints := 0
	// Punkte mit mehr als einer Linie
	overlappedPoints := 0

	for _, n := range overlaps {
		if n > 1 {
			overlappedPoints++
		} else {
			singlePoints++
		}
	}

	assert.Equal(t, 16, singlePoints)
	assert.Equal(t, 5, overlappedPoints)
}

func TestGetOverlapsDiag(t *testing.T) {
	strs, err := ReadLines("day05_test.txt")
	assert.NoError(t, err)

	lines := ParseLines(strs)
	overlaps := GetOverlaps(lines, true)

	// Punkte mit mehr als einer Linie
	overlappedPoints := 0

	for _, n := range overlaps {
		if n > 1 {
			overlappedPoints++
		}
	}

	assert.Equal(t, 12, overlappedPoints)
}