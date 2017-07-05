package sudoku

import (
	"fmt"
	"reflect"
	"testing"
)

var terminalJson = []byte(`
{
	"E":2,
	"C":[
    {
      "B":0,
      "D":1
    },
    {
      "B":0,
      "D":2
    },
    {
      "B":1,
      "D":1
    },
    {
      "B":1,
      "D":0
    }
  ]
}
`)

func Test_Raw2Terminal(t *testing.T) {
	terminal, err := Raw2TerminalJson(terminalJson)
	if err != nil {
		t.Error(err)
	}
	if terminal.E != 2 || terminal.C[3].B != 1 || terminal.C[3].D != 0 {
		t.Error("Incorrect puzzle data")
	}
}

func Test_Terminal2Raw(t *testing.T) {
	terminal, _ := Raw2TerminalJson(terminalJson)
	terminalJson2, _ := TerminalJson2Raw(terminal)
	terminal2, _ := Raw2TerminalJson(terminalJson2)
	if !reflect.DeepEqual(terminal2, terminal) {
		t.Error("reflect.DeepEqual(terminal2, terminal) should be true")
	}
}

func Test_Clone(t *testing.T) {
	rawT, _ := Raw2TerminalJson(terminalJson)
	copyT := rawT.Clone()
	copyT.E = 9
	if rawT.E != 2 || copyT.E != 9 {
		t.Error("Clone failed")
	}
}

func Test_newTerminal(t *testing.T) {
	fmt.Printf("Test_newTerminal:\n%v\n", NewTerminalJson(9))
}
