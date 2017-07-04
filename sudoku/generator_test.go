package sudoku

import (
	"fmt"
	"testing"
)

func Test_GenByte_NORMAL(t *testing.T) {
	fmt.Printf("Test_GenByte_NORMAL:\n%v\n", genTerminal(9, NORMAL, 7, 77))
}

func Test_GenByte_RAMDON(t *testing.T) {
	fmt.Printf("Test_GenByte_RAMDON:\n%v\n", genTerminal(9, RANDOM, 4, 44))
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	digitsDisorder(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
