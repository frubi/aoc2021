package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadDraws(t *testing.T) {
	lines, err := ReadLines("day04_test.txt")
	assert.NoError(t, err)

	exp := []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}

	draws := ReadDraws(lines[0])
	assert.ElementsMatch(t, exp, draws)
}


func TestReadBoard(t *testing.T) {
	lines, err := ReadLines("day04_test.txt")
	assert.NoError(t, err)

	exp := Board{
		[5]int{22, 13, 17, 11, 0},
		[5]int{8, 2, 23, 4, 24},
		[5]int{21, 9, 14, 16, 7},
		[5]int{6, 10, 3, 18, 5},
		[5]int{1, 12, 20, 15, 19},
	}

	board := ReadBoard(lines[1:])
	assert.ElementsMatch(t, exp, board)
}

func TestCheckBingo(t *testing.T) {
	marks1 := Board{
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
	}
	assert.False(t, CheckBingo(marks1))

	marks2 := Board{
		[5]int{1, 1, 1, 1, 0},
		[5]int{1, 1, 0, 1, 1},
		[5]int{1, 0, 1, 0, 1},
		[5]int{1, 0, 0, 0, 1},
		[5]int{0, 1, 1, 1, 0},
	}
	assert.False(t, CheckBingo(marks2))

	marks3 := Board{
		[5]int{0, 0, 0, 0, 1},
		[5]int{0, 0, 0, 0, 1},
		[5]int{0, 0, 0, 0, 1},
		[5]int{0, 0, 0, 0, 1},
		[5]int{0, 0, 0, 0, 1},
	}
	assert.True(t, CheckBingo(marks3))

	marks4 := Board{
		[5]int{1, 0, 0, 0, 0},
		[5]int{1, 0, 0, 0, 0},
		[5]int{1, 0, 0, 0, 0},
		[5]int{1, 0, 0, 0, 0},
		[5]int{1, 0, 0, 0, 0},
	}
	assert.True(t, CheckBingo(marks4))

	marks5 := Board{
		[5]int{1, 1, 1, 1, 1},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
		[5]int{0, 0, 0, 0, 0},
	}
	assert.True(t, CheckBingo(marks5))
}

func TestRunBingoWin(t *testing.T) {
	lines, err := ReadLines("day04_test.txt")
	assert.NoError(t, err)

	boards, draws := ReadGame(lines)
	score := RunBingo(boards, draws, false)

	assert.Equal(t, 4512, score)
}

func TestRunBingoLoose(t *testing.T) {
	lines, err := ReadLines("day04_test.txt")
	assert.NoError(t, err)

	boards, draws := ReadGame(lines)
	score := RunBingo(boards, draws, true)

	assert.Equal(t, 1924, score)
}