package sudoku

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Raw2Terminal(t *testing.T) {
	terminal, err := Raw2TerminalJson(terminalJson_4x4_3)
	if err != nil {
		t.Error(err)
	}
	if terminal.E != 4 || terminal.C[2].B != 1 || terminal.C[2].D != 0 {
		t.Error("Incorrect puzzle data")
	}
}

func Test_Terminal2Raw(t *testing.T) {
	terminal, _ := Raw2TerminalJson(terminalJson_4x4_3)
	terminalJson2, _ := TerminalJson2Raw(terminal)
	terminal2, _ := Raw2TerminalJson(terminalJson2)
	if !reflect.DeepEqual(terminal2, terminal) {
		t.Error("reflect.DeepEqual(terminal2, terminal) should be true")
	}
}

func Test_Clone(t *testing.T) {
	rawT, _ := Raw2TerminalJson(terminalJson_4x4_3)
	copyT := rawT.Clone()
	copyT.E = 44
	if rawT.E != 4 || copyT.E != 44 {
		t.Error("Clone failed")
	}
}

func Test_newTerminal(t *testing.T) {
	fmt.Printf("Test_newTerminal:\n%v\n", NewTerminalJson(9))
}
