package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestReadLines(t *testing.T) {
	expect := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	
	out, err := ReadLines("day02_test.txt")
	assert.NoError(t, err)
	assert.ElementsMatch(t, expect, out)
}


func TestParseCommands(t *testing.T) {
	lines := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	commands := Commands{
		&Command{"forward", 5},
		&Command{"down", 5},
		&Command{"forward", 8},
		&Command{"up", 3},
		&Command{"down", 8},
		&Command{"forward", 2},
	}
	
	out, err := ParseCommands(lines)
	assert.NoError(t, err)
	assert.ElementsMatch(t, commands, out)
}

func TestRunCommands(t *testing.T) {
	state := new(State)

	commands := Commands{
		&Command{"forward", 5},
		&Command{"down", 5},
		&Command{"forward", 8},
		&Command{"up", 3},
		&Command{"down", 8},
		&Command{"forward", 2},
	}

	state.RunCommands(commands)

	assert.Equal(t, 15, state.Horizontal)
	assert.Equal(t, 10, state.Depth)
}

func TestRunAimedCommands(t *testing.T) {
	state := new(State)

	commands := Commands{
		&Command{"forward", 5},
		&Command{"down", 5},
		&Command{"forward", 8},
		&Command{"up", 3},
		&Command{"down", 8},
		&Command{"forward", 2},
	}

	state.RunAimedCommands(commands)

	assert.Equal(t, 15, state.Horizontal)
	assert.Equal(t, 60, state.Depth)
}
