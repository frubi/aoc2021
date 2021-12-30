package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func Sum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}


func ReadNumbers(fn string) ([]int, error) {
	raw, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(raw), "\n")

	numbers := make([]int, 0, len(lines))

	for _, token := range lines {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		value, err := strconv.Atoi(token)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, value)
	}

	return numbers, nil
}

func DepthIncreases(measurements []int) int {
	increases := 0

	last := measurements[0]
	for _, m := range measurements {
		if m > last {
			increases++
		}

		last = m
	}

	return increases
}

func DepthIncreasesWindowed(measurements []int) int {
	increases := 0

	window := 3

	lasts := make([]int, window)
	copy(lasts, measurements)

	for _, m := range measurements[window:] {
		this := make([]int, window)
		copy(this, lasts[1:])
		this[window-1] = m

		if Sum(this) > Sum(lasts) {
			increases++
		}

		lasts = this
	}

	return increases
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "MEASUREMENTS")
		return
	}

	measurements, err := ReadNumbers(os.Args[1])
	if err != nil {
		fmt.Printf("ReadNumbers(%q) failed: %v\n", os.Args[1], err)
		return
	}


	increases := DepthIncreases(measurements)
	fmt.Printf("DepthIncreases: %d\n", increases)

	increases = DepthIncreasesWindowed(measurements)
	fmt.Printf("DepthIncreasesWindowed: %d\n", increases)
}
