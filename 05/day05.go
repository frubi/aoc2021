package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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

func ParsePoint(str string) Point {
	p := Point{}

	tokens := strings.Split(str, ",")
	p.X = SafeAtoi(tokens[0])
	p.Y = SafeAtoi(tokens[1])

	return p
}

func ParseLine(str string) Line {
	l := Line{}

	tokens := strings.Split(str, " -> ")
	l.Start = ParsePoint(tokens[0])
	l.End = ParsePoint(tokens[1])

	return l
}

func ParseLines(strs []string) []Line {
	lines := make([]Line, 0, len(strs))

	for _, str := range strs {
		l := ParseLine(str)
		lines = append(lines, l)
	}

	return lines
}

func PointsEqual(a Point, b Point) bool {
	if a.X != b.X {
		return false
	}

	if a.Y != b.Y {
		return false
	}

	return true
}

func GetBounds(lines []Line) (Point, Point) {
	// Minimum und Maximum initialisieren
	min := lines[0].Start
	max := lines[0].End

	for _, line := range lines {
		min.X = Min(min.X, line.Start.X)
		min.X = Min(min.X, line.End.X)
		min.Y = Min(min.Y, line.Start.Y)
		min.Y = Min(min.Y, line.End.Y)

		max.X = Max(max.X, line.Start.X)
		max.X = Max(max.X, line.End.X)
		max.Y = Max(max.Y, line.Start.Y)
		max.Y = Max(max.Y, line.End.Y)
	}

	return min, max
}

func GetLinePoints(line Line, includeDiag bool) []Point {
	points := make([]Point, 0)

	dx := 0
	dy := 0

	// Änderung Y
	if line.Start.Y < line.End.Y {
		dy = 1
	}
	if line.Start.Y > line.End.Y {
		dy = -1
	}

	// Änderung X
	if line.Start.X < line.End.X {
		dx = 1
	}
	if line.Start.X > line.End.X {
		dx = -1
	}

	if (dx == 0) && (dy == 0) {
		return points
	}

	if (includeDiag == false) && (dx != 0) && (dy != 0) {
		return points
	}

	here := line.Start
	for {
		points = append(points, here)

		if PointsEqual(line.End, here) {
			break
		}

		here.X += dx
		here.Y += dy
	}

	return points
}

func GetOverlaps(lines []Line, includeDiag bool) map[Point]int {
	overlaps := make(map[Point]int)

	for _, line := range lines {
		points := GetLinePoints(line, includeDiag)

		for _, point := range points {
			n, known := overlaps[point]
			if known {
				overlaps[point] = n + 1
			} else {
				overlaps[point] = 1
			}
		}
	}

	return overlaps
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "LINES")
		return
	}

	strs, err := ReadLines(os.Args[1])
	if err != nil {
		fmt.Println("ReadLines() failed:", err)
		return
	}

	lines := ParseLines(strs)


	fmt.Println("=== horizontal / vertical ===")
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

	fmt.Println("Single points:", singlePoints)
	fmt.Println("Overlapped points:", overlappedPoints)


	fmt.Println("=== horizontal / vertical / diagonal ===")
	overlaps = GetOverlaps(lines, true)

	// Punkte mit nur einer Linie
	singlePoints = 0
	// Punkte mit mehr als einer Linie
	overlappedPoints = 0

	for _, n := range overlaps {
		if n > 1 {
			overlappedPoints++
		} else {
			singlePoints++
		}
	}

	fmt.Println("Single points:", singlePoints)
	fmt.Println("Overlapped points:", overlappedPoints)

}
