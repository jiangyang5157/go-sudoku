package sudoku

import (
	"fmt"
	"testing"
)

func Test_genMaterial(t *testing.T) {
	fmt.Printf("Test_genMaterial:\n%v\n", NewTerminalJson(9).genMaterial())
}

func Test_genBlock(t *testing.T) {
	fmt.Printf("Test_genBlock SQUARE:\n%v\n", NewTerminalJson(9).genBlock(SQUARE))
	fmt.Printf("Test_genBlock RAMDOM:\n%v\n", NewTerminalJson(9).genBlock(RANDOM))
	fmt.Printf("Test_genBlock IRREGULAR:\n%v\n", NewTerminalJson(9).genBlock(IRREGULAR))
}

func Test_reachableCells(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := newGraph(terminal)

	index := 0
	fmt.Printf("Test_reachableCells %d: %d\n", index, reachableCells(g, index2id(index)))
	neighbours := srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, index2id(neighbour), index2id(index))
	}

	index = 1
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, index2id(neighbour), index2id(index))
	}
	fmt.Printf("Test_reachableCells %d: %d\n", 0, reachableCells(g, index2id(0)))

	index = 10
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, index2id(neighbour), index2id(index))
	}
	fmt.Printf("Test_reachableCells %d: %d\n", 1, reachableCells(g, index2id(1)))

	index = 19
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, index2id(neighbour), index2id(index))
	}
	fmt.Printf("Test_reachableCells %d: %d\n", 10, reachableCells(g, index2id(10)))

	index = 18
	neighbours = srcNeighbours(terminal, g, index)
	for _, neighbour := range neighbours {
		unlink(g, index2id(neighbour), index2id(index))
	}
	fmt.Printf("Test_reachableCells %d: %d\n", 19, reachableCells(g, index2id(19)))
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
