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

func Test_increasingDigits_digitsDisorder(t *testing.T) {
	var digits []int = increasingDigits(1, 10)
	for i := 0; i < 10; i++ {
		if digits[i] != i+1 {
			t.Error("increasingDigits is wrong")
		}
	}
	disorderDigits(digits)
	ok := false
	for i := 0; i < 10; i++ {
		if digits[i] != i+1 {
			ok = true
			break
		}
	}
	if !ok {
		t.Error("disorderDigits is wrong")
	}
}
