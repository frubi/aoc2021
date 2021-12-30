package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func SafeAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid number %q: %v", s, err))
	}

	return v
}

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

func ParseNumbers(str string) []int {
	tokens := strings.Split(str, ",")

	numbers := make([]int, len(tokens))
	for index, token := range tokens {
		numbers[index] = SafeAtoi(token)
	}

	return numbers
}

func LanternfishCycle(in []int) []int {
	out := make([]int, len(in))

	for index, timer := range in {
		if timer == 0 {
			out[index] = 6
			out = append(out, 8)
		} else {
			out[index] = timer - 1
		}
	}

	return out
}

func LanternfishCycleN(in []int, days int) []int {
	state := make([]int, len(in))
	copy(state, in)

	for day := 0; day < days; day++ {
		state = LanternfishCycle(state)
	}

	return state
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "LIST DAYS")
		return
	}

	lines, err := ReadLines(os.Args[1])
	if err != nil {
		fmt.Println("ReadLines() failed:", err)
		return
	}

	days := SafeAtoi(os.Args[2])
	start := ParseNumbers(lines[0])

	out := LanternfishCycleN(start, days)
	fmt.Println("Number of laternfish after", days, "days:", len(out))

}
