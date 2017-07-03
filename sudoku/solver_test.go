package sudoku

import (
	"fmt"
	"strings"
	"testing"
)

func Test_SolvePuzzleByRaw(t *testing.T) {
	raw :=
		// "......123" +
		// 	"..9......" +
		// 	".....9..." +
		// 	"........." +
		// 	"........." +
		// 	"........." +
		// 	"........." +
		// 	"........." +
		// 	"........." // 0 solutions puzzle

		// "........." +
		// 	"..41.26.." +
		// 	".3..5..2." +
		// 	".2..1..3." +
		// 	"..65.41.." +
		// 	".8..7..4." +
		// 	".7..2..6." +
		// 	"..14.35.." +
		// 	"........." // 1 solutions puzzle

		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

		// "....7.94." +
		// 	".7..9...5" +
		// 	"3....5.7." +
		// 	"..74..1.." +
		// 	"463.8...." +
		// 	".....7.8." +
		// 	"8..7....." +
		// 	"7......28" +
		// 	".5..68..." // 188 solutions puzzle

		// "." // 1 solutions puzzle

		// "...." +
		// 	".4.." +
		// 	"2..." +
		// 	"..43" // 0 solutions puzzle

		// "...." +
		// 	".4.." +
		// 	"2..." +
		// 	"...3" // 3 solutions puzzle

		// "................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" +
		// 	"................" // ? solutions puzzle

	squares := 3
	printPuzzleRaw(squares, raw)
	solutions := 200
	sols := strings.Split(SolveRaw(squares, raw, solutions), string(SOLUTION_PREFIX))
	fmt.Printf("Searching for %d solutions, %d solutions found.\n", solutions, len(sols)-1)
	for i := 1; i < len(sols); i++ {
		printPuzzleRaw(squares, sols[i])
	}
}

func Test_solve(t *testing.T) {
	raw :=
		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

	squares := 3
	p := newPuzzle(squares)
	p.build(raw2digits(raw))
	solutions := 2
	sols := strings.Split(p.solve(solutions), string(SOLUTION_PREFIX))
	fmt.Printf("Searching for %d solutions, %d solutions found.\n", solutions, len(sols)-1)
	for i := 1; i < len(sols); i++ {
		printPuzzleRaw(squares, sols[i])
	}

	raw =
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..." // 188 solutions puzzle

	p.build(raw2digits(raw))
	solutions = 1
	sols = strings.Split(p.solve(solutions), string(SOLUTION_PREFIX))
	fmt.Printf("Searching for %d solutions, %d solutions found.\n", solutions, len(sols)-1)
	for i := 1; i < len(sols); i++ {
		printPuzzleRaw(squares, sols[i])
	}
}

func Test_HasUniqueSolution(t *testing.T) {
	p := newPuzzle(3)

	raw :=
		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

	p.build(raw2digits(raw))
	if p.HasUniqueSolution() {
		t.Error("This is 2 solutions puzzle.")
	}

	raw =
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..." // 188 solutions puzzle

	p.build(raw2digits(raw))
	if p.HasUniqueSolution() {
		t.Error("This is 188 solutions puzzle.")
	}

	raw =
		"......123" +
			"..9......" +
			".....9..." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........." // 0 solutions puzzle

	p.build(raw2digits(raw))
	if p.HasUniqueSolution() {
		t.Error("This is 0 solutions puzzle.")
	}

	raw =
		"........." +
			"..41.26.." +
			".3..5..2." +
			".2..1..3." +
			"..65.41.." +
			".8..7..4." +
			".7..2..6." +
			"..14.35.." +
			"........." // 1 solutions puzzle

	p.build(raw2digits(raw))
	if !p.HasUniqueSolution() {
		t.Error("This is 1 solutions puzzle.")
	}
}
