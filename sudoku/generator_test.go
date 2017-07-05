package sudoku

import (
	"fmt"
	"testing"
)

func Test_genTerminal(t *testing.T) {
	fmt.Printf("Test_genTerminal_9x9:\n%v\n", genTerminal(9, SQUARE, 4, 44))
}

func Test_genMaterial(t *testing.T) {
	fmt.Printf("Test_genMaterial:\n%v\n", NewTerminal(9).genMaterial())
}

func Test_genBlock(t *testing.T) {
	fmt.Printf("Test_genBlock SQUARE:\n%v\n", NewTerminal(9).genBlock(SQUARE))
	fmt.Printf("Test_genBlock RAMDOM:\n%v\n", NewTerminal(9).genBlock(RANDOM))
	fmt.Printf("Test_genBlock IRREGULAR:\n%v\n", NewTerminal(9).genBlock(IRREGULAR))
}

func Test_genPuzzle(t *testing.T) {
	// TODO
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
