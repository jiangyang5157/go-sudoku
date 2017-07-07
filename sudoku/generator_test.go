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

func Test_visiteUnblockedCells(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)
	reachable := visiteUnblockedCells(g, index2id(9))
	if reachable != 81 {
		t.Error()
	}
	unlinkFromUnblocked(terminal, g, 0)
	unlinkFromUnblocked(terminal, g, 1)
	unlinkFromUnblocked(terminal, g, 10)
	unlinkFromUnblocked(terminal, g, 19)

	unlinkFromUnblocked(terminal, g, 18)
	reachable = visiteUnblockedCells(g, index2id(9))
	if reachable != 1 {
		t.Error()
	}
	linkFromUnblocked(terminal, g, 18)
	reachable = visiteUnblockedCells(g, index2id(9))
	if reachable != 77 {
		t.Error()
	}

	terminal.C[0].B = 1
	terminal.C[1].B = 1
	terminal.C[10].B = 1
	terminal.C[19].B = 1
	terminal.C[18].B = 1
	unlinkFromUnblocked(terminal, g, 18)
	unlink(terminal, g, 19, 18)
	reachable = visiteUnblockedCells(g, index2id(19))
	if reachable != 75 {
		t.Error()
	}
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
