package sudoku

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GeneratePuzzle(t *testing.T) {
	squares := 4
	minSubGiven := 10
	minTotalGiven := 80
	raw := GeneratePuzzle(squares, minSubGiven, minTotalGiven)
	printPuzzleRaw(squares, raw)
	p := newPuzzleData(squares)
	p.build(raw2digits(raw))
	solutions := 2
	sols := strings.Split(p.solve(solutions), string(SOLUTION_PREFIX))
	fmt.Printf("Searching for %d solutions, %d solutions found.\n", solutions, len(sols)-1)
	for i := 1; i < len(sols); i++ {
		printPuzzleRaw(squares, sols[i])
	}
}

func Test_generateTerminalPuzzle(t *testing.T) {
	squares := 3
	p := newPuzzleData(squares)
	tp := p.generateTerminalPuzzle()
	printPuzzleDigits(squares, tp)
}
