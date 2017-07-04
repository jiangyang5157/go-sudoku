package sudoku

import "testing"

func Test_Solve_rawJson(t *testing.T) {
	terminal := SolveByte(terminalJson_9x9_0)
	if string(terminal) != "null" {
		t.Error("terminalJson_9x9_0 shouldn't have solution")
	}
	terminal = SolveByte(terminalJson_9x9_2)
	if string(terminal) == "null" {
		t.Error("terminalJson_9x9_2 should have solution")
	}
	terminal = SolveByte(terminalJson_1x1_1)
	if string(terminal) == "null" {
		t.Error("terminalJson_1x1_1 should have solution")
	}
	terminal = SolveByte(terminalJson_4x4_0)
	if string(terminal) != "null" {
		t.Error("terminalJson_4x4_0 shouldn't have solution")
	}
	terminal = SolveByte(terminalJson_4x4_3)
	if string(terminal) == "null" {
		t.Error("terminalJson_4x4_3 should have solution")
	}
}
