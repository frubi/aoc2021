package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [5][5]int

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

func ReadDraws(line string) []int {
	draws := make([]int, 0)

	tokens := strings.Split(line, ",")
	for _, token := range tokens {
		v := SafeAtoi(token)
		draws = append(draws, v)
	}

	return draws
}

func ReadBoard(lines []string) Board {
	board := Board{}

	if len(lines) < 5 {
		panic(fmt.Sprintf("too few lines left: %d", len(lines)))
	}

	for row := 0; row < 5; row++ {
		tokens := strings.Split(lines[row], " ")

		col := 0
		for _, token := range tokens {
			if token == "" {
				continue
			} 

			board[row][col] = SafeAtoi(token)
			col++
		}
	}

	return board
}

func ReadGame(lines []string) ([]Board, []int) {
	draws := ReadDraws(lines[0])

	boards := make([]Board, 0)

	for offset := 1; offset < len(lines); offset += 5 {
		board := ReadBoard(lines[offset:])
		boards = append(boards, board)
	}

	return boards, draws
}

func CheckBingo(marks Board) bool {
	// Reihen prüfen
	for row := 0; row < 5; row++ {
		isBingo := true
		for col := 0; col < 5; col++ {
			if marks[row][col] != 1 {
				isBingo = false
			}
		}

		if isBingo {
			return true
		}
	}

	// Spalten prüfen
	for col := 0; col < 5; col++ {
		isBingo := true
		for row := 0; row < 5; row++ {
			if marks[row][col] != 1 {
				isBingo = false
			}
		}

		if isBingo {
			return true
		}
	}

	return false
}


func RunBingo(boards []Board, draws []int, untilLast bool) int {
	// Markierung
	// 0 -> Zahl wurde noch nicht gezogen
	// 1 -> Zahl wurde gezogen
	marks := make([]Board, len(boards))

	// Kennzeichen, ob das Board bereits abgeschlossen ist
	completed := make(map[int]bool)

	for _, draw := range draws {
		for index, board := range boards {
			// Bereits abgeschlossene Boards nicht weiter berücksichtigen
			if _, isCompleted := completed[index]; isCompleted {
				continue
			}

			// Flag für Änderungen an den Markierungen des Boards
			marksChanged := false

			// Prüfen, ob die gezogene Zahl auf dem Board ist
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if board[row][col] == draw {
						marks[index][row][col] = 1
						marksChanged = true
					}
				}
			}

			// Prüfen, ob ein Bingo vorliegt
			if marksChanged && CheckBingo(marks[index]) {
				// Markierung für erfolgreichen Abschluss setzen
				completed[index] = true

				// Summe der Zahlen auf dem Gewinner-Board, die nicht
				// markiert (aka nicht gezogen) sind.
				sum := 0
				for row := 0; row < 5; row++ {
					for col := 0; col < 5; col++ {
						if marks[index][row][col] == 0 {
							sum += board[row][col]
						}
					}
				}

				// Ergebnis aus Summe und letzter Ziehung ermitteln
				score := sum * draw

				if untilLast == true {
					if len(boards) == len(completed) {
						return score
					}
				} else {
					return score
				}
			}
		}
	}

	// Abgeschlossen ohne Gewinner
	return -1
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args, "GAME")
		return
	}

	lines, err := ReadLines(os.Args[1])
	if err != nil {
		fmt.Println("ReadLines() failed:", err)
		return
	}

	boards, draws := ReadGame(lines)
	
	score := RunBingo(boards, draws, false)
	fmt.Println("Winning score:", score)

	score = RunBingo(boards, draws, true)
	fmt.Println("Loosing score:", score)
}
