package sudoku

import (
	"fmt"
	"testing"
)

func Test_newSudoku(t *testing.T) {
	terminal, _ := Raw2Terminal(terminalJson)
	s := newSudoku(terminal)
	fmt.Printf("%v\n", s.t)
}

func Test_initialize(t *testing.T) {
	terminal, _ := Raw2Terminal(terminalJson)
	s := newSudoku(terminal)
	s.initialize()
	fmt.Printf("%v\n", s.t)
}
