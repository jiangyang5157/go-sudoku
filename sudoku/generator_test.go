package sudoku

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GeneratePuzzle(t *testing.T) {
	squares := 3
	minSubGiven := 2
	minTotalGiven := 17
	raw := GeneratePuzzle(squares, minSubGiven, minTotalGiven)
	p := newPuzzle(squares)
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
	p := newPuzzle(squares)
	printPuzzleDigits(squares, p.generateTerminalPuzzle())
}
