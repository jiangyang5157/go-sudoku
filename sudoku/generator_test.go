package sudoku

import (
	"fmt"
	"testing"
)

func Test_genTerminal(t *testing.T) {
	fmt.Printf("Test_genTerminal_9x9:\n%v\n", genTerminal(9, REGULAR, 4, 44))
}

func Test_genMaterial(t *testing.T) {
	fmt.Printf("Test_genMaterial:\n%v\n", newTerminal(9).genMaterial())
}

func Test_genBlock_REGULAR(t *testing.T) {
	fmt.Printf("Test_genBlock_REGULAR:\n%v\n", newTerminal(9).genBlock(REGULAR))
}

func Test_genBlock_IRREGULAR(t *testing.T) {
	fmt.Printf("Test_genBlock_IRREGULAR:\n%v\n", newTerminal(9).genBlock(IRREGULAR))
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	digitsDisorder(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
