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
	// fmt.Printf("Test_genBlock IRREGULAR:\n%v\n", NewTerminalJson(9).genBlock(IRREGULAR))
}

// func Test_visiteUnblockedCells(t *testing.T) {
// 	terminal := NewTerminalJson(9)
// 	g := NewGraph(terminal)
// 	reachable := reachableCells(g, index2id(0))
// 	if reachable != 81 {
// 		t.Error()
// 	}
// 	unlinkedByLinkedNeighbours(terminal, g, 0)
// 	unlinkedByLinkedNeighbours(terminal, g, 1)
// 	unlinkedByLinkedNeighbours(terminal, g, 10)
// 	unlinkedByLinkedNeighbours(terminal, g, 19)
// 	reachable = reachableCells(g, index2id(2))
// 	if reachable != 77 {
// 		t.Error()
// 	}
// 	unlinkedByLinkedNeighbours(terminal, g, 18)
// 	reachable = reachableCells(g, index2id(2))
// 	if reachable != 75 {
// 		t.Error()
// 	}
// }

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
