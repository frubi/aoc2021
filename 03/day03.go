package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func SafeBinToDec(s string) int {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid number %q: %v", s, err))
	}

	return int(ui)
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

func MostCommonBit(lines []string, offset int) int {
	stats := []int{ 0, 0 }

	for _, line := range lines {
		char := line[offset]

		switch char {
		case '0':
			stats[0]++
		case '1':
			stats[1]++
		default:
			panic(fmt.Sprintf("invalid char: %c" , char))
		}
	}

	if stats[1] >= stats[0] {
		return 1
	} else {
		return 0
	}
}

func LeastCommonBit(lines []string, offset int) int {
	msc := MostCommonBit(lines, offset)

	if msc == 1 {
		return 0
	} else {
		return 1
	}
}

type ReportFunc func([]string, int) int

func BuildReport(lines []string, repFn ReportFunc) int {
	output := 0
	length := len(lines[0])

	for offset := 0; offset < length; offset++ {
		v := repFn(lines, offset)

		output = output << 1
		if v == 1 {
			output |= 1
		}
	}

	return output
}

func FilterReport(lines []string, repFn ReportFunc, offset int) string {
	v := repFn(lines, offset)

	newLines := make([]string, 0)
	for _, line := range lines {
		keepLine := false

		if (v == 1) && (line[offset] == '1') {
			keepLine = true
		}
		if (v == 0) && (line[offset] == '0') {
			keepLine = true
		}

		if keepLine == true {
			newLines = append(newLines, line)
		}
	}

	if len(newLines) == 1 {
		return newLines[0]
	} else {
		return FilterReport(newLines, repFn, offset + 1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "DIAG")
		return
	}

	lines, err := ReadLines(os.Args[1])
	if err != nil {
		fmt.Println("ReadLines() failed", err)
		return
	}

	gamma := BuildReport(lines, MostCommonBit)
	fmt.Printf("gamma = [base2]%b [base10]%d\n", gamma, gamma)

	epsilon := BuildReport(lines, LeastCommonBit)
	fmt.Printf("epsilon = [base2]%b [base10]%d\n", epsilon, epsilon)

	power := gamma * epsilon
	fmt.Printf("power = %d\n", power)


	oxygenStr := FilterReport(lines, MostCommonBit, 0)
	oxygen := SafeBinToDec(oxygenStr)
	fmt.Printf("oxygen = %d (from %q)\n", oxygen, oxygenStr)

	co2Str := FilterReport(lines, LeastCommonBit, 0)
	co2 := SafeBinToDec(co2Str)
	fmt.Printf("co2 = %d (from %q)\n", co2, co2Str)

	lsr := oxygen * co2
	fmt.Printf("life support rating = %d\n", lsr)
}