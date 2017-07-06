package sudoku

import (
	"fmt"
	"testing"
)

func Test_NewGraph(t *testing.T) {
	tJson, _ := Raw2TerminalJson(terminalJson_4x4_3)
	fmt.Printf("Test_NewGraph Raw2TerminalJson\n%v\n", tJson)
	NewGraph(tJson)
	g := NewGraph(tJson)
	fmt.Printf("Test_NewGraph ToGraph\n%v\n", g)
}
