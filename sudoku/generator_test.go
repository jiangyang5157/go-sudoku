package sudoku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_genBlock(t *testing.T) {
	fmt.Printf("Test_genBlock:\n%v\n", NewTerminalJson(9).genBlock())
}

func Test_GenGenTerminalJson(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Test_GenTerminalJson:\n%v\n", GenTerminalJson(9, 2, 22))
}

func Test_GenString(t *testing.T) {
	fmt.Printf("Test_GenString:\n%v\n", GenString(9, 2, 22))
}

func Test_increasingDigits_digitsDisorder(t *testing.T) {
	var digits []int = increasingDigits(1, 10)
	fmt.Printf("Test_increasingDigits:\n%d\n", digits)
	disorderDigits(digits)
	fmt.Printf("Test_digitsDisorder:\n%d\n", digits)
}
