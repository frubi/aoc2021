package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type State struct {
	Horizontal int
	Depth int
	Aim int
}

type Command struct {
	Direction string
	Units int
}

type Commands []*Command


func ReadLines(fn string) ([]string, error) {
	raw, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(raw), "\n")

	out := make([]string, 0, len(lines))

	for _, token := range lines {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		out = append(out, token)
	}

	return out, nil
}

func ParseCommands(lines []string) (Commands, error) {
	commands := make(Commands, 0)

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid number of tokens: %q", line)
		}

		units, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, fmt.Errorf("Invalid units value: %v", err)
		}

		cmd := new(Command)
		cmd.Direction = tokens[0]
		cmd.Units = units

		commands = append(commands, cmd)
	}

	return commands, nil
}

func (s *State) RunCommands(cmds Commands) {
	for _, cmd := range cmds {
		switch cmd.Direction {
		case "forward":
			s.Horizontal += cmd.Units
		case "down":
			s.Depth += cmd.Units
		case "up":
			s.Depth -= cmd.Units
		}
	}
}

func (s *State) RunAimedCommands(cmds Commands) {
	for _, cmd := range cmds {
		switch cmd.Direction {
		case "forward":
			s.Horizontal += cmd.Units
			s.Depth += cmd.Units * s.Aim
		case "down":
			s.Aim += cmd.Units
		case "up":
			s.Aim -= cmd.Units
		}
	}
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "COMMANDS")
		return
	}

	lines, err := ReadLines(os.Args[1])
	if err != nil {
		fmt.Println("ReadLines() failed:", err)
		return
	}

	commands, err := ParseCommands(lines)
	if err != nil {
		fmt.Println("ParseCommands() failed:", err)
		return
	}


	fmt.Println("=== RunCommands ===")

	state := new(State)
	state.RunCommands(commands)

	fmt.Printf("Horizontal: %d\n", state.Horizontal)
	fmt.Printf("Depth: %d\n", state.Depth)
	fmt.Printf("Horizontal * Depth: %d\n", state.Horizontal * state.Depth)


	fmt.Println("=== RunAimedCommands ===")

	state = new(State)
	state.RunAimedCommands(commands)

	fmt.Printf("Horizontal: %d\n", state.Horizontal)
	fmt.Printf("Depth: %d\n", state.Depth)
	fmt.Printf("Horizontal * Depth: %d\n", state.Horizontal * state.Depth)

}