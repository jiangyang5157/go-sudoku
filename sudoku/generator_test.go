package sudoku

import (
	"fmt"
	"testing"
)

func Test_genBlock(t *testing.T) {
	fmt.Printf("Test_genBlock SQUARE:\n%v\n", NewTerminalJson(9).genBlock(SQUARE))
}

func Test_GenTerminal_genPuzzle(t *testing.T) {
	tSQUARE := GenTerminal(9, SQUARE)
	fmt.Printf("Test_GenTerminal SQUARE:\n%v\n", tSQUARE)
	fmt.Printf("Test_genPuzzle SQUARE:\n%v\n", tSQUARE.genPuzzle(5, 55))
}

func Test_GenString(t *testing.T) {
	fmt.Printf("Test_GenString SQUARE:\n%v\n", GenString(9, 0, 5, 55))
}

func Test_reachableNum(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)

	var begin int
	index := 0
	neighbours := srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 1
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 10
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 19
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 18
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))
	for _, neighbour := range neighbours {
		link(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 28
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 27
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))
	for _, neighbour := range neighbours {
		link(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 37
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 38
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 29
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 20
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 2
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 11
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))

	index = 12
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, Index2Id(neighbour), Index2Id(index))
	}
	for i := len(terminal.C) - 1; i >= 0; i-- {
		if terminal.C[i].B == 0 {
			begin = i
			break
		}
	}
	fmt.Printf("Test_reachableNum %d: %d\n", begin, reachableNum(g, Index2Id(begin)))
	// simulate 0 1 10 19 (18) 28 (27) 37 38 29 20, 2 11 12
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
