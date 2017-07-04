package sudoku

import (
	"reflect"
	"testing"
)

var terminalJson = []byte(`
{
	"E":2,
	"C":[
    {
      "I":0,
      "J":0,
      "B":0,
      "D":1
    },
    {
      "I":0,
      "J":1,
      "B":0,
      "D":2
    },
    {
      "I":1,
      "J":0,
      "B":1,
      "D":1
    },
    {
      "I":1,
      "J":1,
      "B":1,
      "D":0
    }
  ]
}
`)

func Test_Raw2Terminal(t *testing.T) {
	terminal, err := Raw2Terminal(terminalJson)
	if err != nil {
		t.Error(err)
	}
	if terminal.E != 2 || terminal.C[3].I != 1 || terminal.C[3].J != 1 || terminal.C[3].B != 1 || terminal.C[3].D != 0 {
		t.Error("Incorrect puzzle data")
	}
}

func Test_Terminal2Raw(t *testing.T) {
	terminal, _ := Raw2Terminal(terminalJson)
	terminalJson2, _ := Terminal2Raw(terminal)
	terminal2, _ := Raw2Terminal(terminalJson2)
	if !reflect.DeepEqual(terminal2, terminal) {
		t.Error("reflect.DeepEqual(terminal2, terminal) should be true")
	}
}

func Test_Clone(t *testing.T) {
	rawT, _ := Raw2Terminal(terminalJson)
	copyT := rawT.Clone()
	copyT.E = 9
	if rawT.E != 2 || copyT.E != 9 {
		t.Error("Clone failed")
	}
}
