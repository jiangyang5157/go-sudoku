package sudoku

import "testing"

func Test_newSudoku_initialize(t *testing.T) {
	terminal, _ := Raw2TerminalJson(terminalJson)
	s := newSudoku(terminal)
	s.initialize()
	if s.t.E != 2 || s.t.C[3].I != 1 || s.t.C[3].J != 1 || s.t.C[3].B != 1 || s.t.C[3].D != 0 {
		t.Error("Incorrect sudoku.t data")
	}
}

func Test_hasUniqueSolution(t *testing.T) {
	terminal, _ := Raw2TerminalJson(terminalJson_9x9_0)
	s := newSudoku(terminal)
	s.initialize()
	if s.hasUniqueSolution() {
		t.Error("terminalJson_9x9_0 should not have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_9x9_1)
	s = newSudoku(terminal)
	s.initialize()
	if !s.hasUniqueSolution() {
		t.Error("terminalJson_9x9_1 should have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_9x9_2)
	s = newSudoku(terminal)
	s.initialize()
	if s.hasUniqueSolution() {
		t.Error("terminalJson_9x9_2 should not have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_9x9_188)
	s = newSudoku(terminal)
	s.initialize()
	if s.hasUniqueSolution() {
		t.Error("terminalJson_9x9_188 should not have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_1x1_1)
	s = newSudoku(terminal)
	s.initialize()
	if !s.hasUniqueSolution() {
		t.Error("terminalJson_1x1_1 should have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_4x4_0)
	s = newSudoku(terminal)
	s.initialize()
	if s.hasUniqueSolution() {
		t.Error("terminalJson_4x4_0 should not have unique solution")
	}

	terminal, _ = Raw2TerminalJson(terminalJson_4x4_3)
	s = newSudoku(terminal)
	s.initialize()
	if s.hasUniqueSolution() {
		t.Error("terminalJson_4x4_3 should not have unique solution")
	}
}

// "......123" +
// 	"..9......" +
// 	".....9..." +
// 	"........." +
// 	"........." +
// 	"........." +
// 	"........." +
// 	"........." +
// 	"........." // 0 solutions puzzle
var terminalJson_9x9_0 = []byte(`
	{
		"E":9,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":0,"D":0},
			{"I":0,"J":3,"B":1,"D":0},
			{"I":0,"J":4,"B":1,"D":0},
			{"I":0,"J":5,"B":1,"D":0},
			{"I":0,"J":6,"B":2,"D":1},
			{"I":0,"J":7,"B":2,"D":2},
			{"I":0,"J":8,"B":2,"D":3},

			{"I":1,"J":0,"B":0,"D":0},
			{"I":1,"J":1,"B":0,"D":0},
			{"I":1,"J":2,"B":0,"D":9},
			{"I":1,"J":3,"B":1,"D":0},
			{"I":1,"J":4,"B":1,"D":0},
			{"I":1,"J":5,"B":1,"D":0},
			{"I":1,"J":6,"B":2,"D":0},
			{"I":1,"J":7,"B":2,"D":0},
			{"I":1,"J":8,"B":2,"D":0},

			{"I":2,"J":0,"B":0,"D":0},
			{"I":2,"J":1,"B":0,"D":0},
			{"I":2,"J":2,"B":0,"D":0},
			{"I":2,"J":3,"B":1,"D":0},
			{"I":2,"J":4,"B":1,"D":0},
			{"I":2,"J":5,"B":1,"D":9},
			{"I":2,"J":6,"B":2,"D":0},
			{"I":2,"J":7,"B":2,"D":0},
			{"I":2,"J":8,"B":2,"D":0},

			{"I":3,"J":0,"B":3,"D":0},
			{"I":3,"J":1,"B":3,"D":0},
			{"I":3,"J":2,"B":3,"D":0},
			{"I":3,"J":3,"B":4,"D":0},
			{"I":3,"J":4,"B":4,"D":0},
			{"I":3,"J":5,"B":4,"D":0},
			{"I":3,"J":6,"B":5,"D":0},
			{"I":3,"J":7,"B":5,"D":0},
			{"I":3,"J":8,"B":5,"D":0},

			{"I":4,"J":0,"B":3,"D":0},
			{"I":4,"J":1,"B":3,"D":0},
			{"I":4,"J":2,"B":3,"D":0},
			{"I":4,"J":3,"B":4,"D":0},
			{"I":4,"J":4,"B":4,"D":0},
			{"I":4,"J":5,"B":4,"D":0},
			{"I":4,"J":6,"B":5,"D":0},
			{"I":4,"J":7,"B":5,"D":0},
			{"I":4,"J":8,"B":5,"D":0},

			{"I":5,"J":0,"B":3,"D":0},
			{"I":5,"J":1,"B":3,"D":0},
			{"I":5,"J":2,"B":3,"D":0},
			{"I":5,"J":3,"B":4,"D":0},
			{"I":5,"J":4,"B":4,"D":0},
			{"I":5,"J":5,"B":4,"D":0},
			{"I":5,"J":6,"B":5,"D":0},
			{"I":5,"J":7,"B":5,"D":0},
			{"I":5,"J":8,"B":5,"D":0},

			{"I":6,"J":0,"B":6,"D":0},
			{"I":6,"J":1,"B":6,"D":0},
			{"I":6,"J":2,"B":6,"D":0},
			{"I":6,"J":3,"B":7,"D":0},
			{"I":6,"J":4,"B":7,"D":0},
			{"I":6,"J":5,"B":7,"D":0},
			{"I":6,"J":6,"B":8,"D":0},
			{"I":6,"J":7,"B":8,"D":0},
			{"I":6,"J":8,"B":8,"D":0},

			{"I":7,"J":0,"B":6,"D":0},
			{"I":7,"J":1,"B":6,"D":0},
			{"I":7,"J":2,"B":6,"D":0},
			{"I":7,"J":3,"B":7,"D":0},
			{"I":7,"J":4,"B":7,"D":0},
			{"I":7,"J":5,"B":7,"D":0},
			{"I":7,"J":6,"B":8,"D":0},
			{"I":7,"J":7,"B":8,"D":0},
			{"I":7,"J":8,"B":8,"D":0},

			{"I":8,"J":0,"B":6,"D":0},
			{"I":8,"J":1,"B":6,"D":0},
			{"I":8,"J":2,"B":6,"D":0},
			{"I":8,"J":3,"B":7,"D":0},
			{"I":8,"J":4,"B":7,"D":0},
			{"I":8,"J":5,"B":7,"D":0},
			{"I":8,"J":6,"B":8,"D":0},
			{"I":8,"J":7,"B":8,"D":0},
			{"I":8,"J":8,"B":8,"D":0}
	  ]
	}
`)

// "........." +
// 	"..41.26.." +
// 	".3..5..2." +
// 	".2..1..3." +
// 	"..65.41.." +
// 	".8..7..4." +
// 	".7..2..6." +
// 	"..14.35.." +
// 	"........." // 1 solutions puzzle
var terminalJson_9x9_1 = []byte(`
	{
		"E":9,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":0,"D":0},
			{"I":0,"J":3,"B":1,"D":0},
			{"I":0,"J":4,"B":1,"D":0},
			{"I":0,"J":5,"B":1,"D":0},
			{"I":0,"J":6,"B":2,"D":0},
			{"I":0,"J":7,"B":2,"D":0},
			{"I":0,"J":8,"B":2,"D":0},

			{"I":1,"J":0,"B":0,"D":0},
			{"I":1,"J":1,"B":0,"D":0},
			{"I":1,"J":2,"B":0,"D":4},
			{"I":1,"J":3,"B":1,"D":1},
			{"I":1,"J":4,"B":1,"D":0},
			{"I":1,"J":5,"B":1,"D":2},
			{"I":1,"J":6,"B":2,"D":6},
			{"I":1,"J":7,"B":2,"D":0},
			{"I":1,"J":8,"B":2,"D":0},

			{"I":2,"J":0,"B":0,"D":0},
			{"I":2,"J":1,"B":0,"D":3},
			{"I":2,"J":2,"B":0,"D":0},
			{"I":2,"J":3,"B":1,"D":0},
			{"I":2,"J":4,"B":1,"D":5},
			{"I":2,"J":5,"B":1,"D":0},
			{"I":2,"J":6,"B":2,"D":0},
			{"I":2,"J":7,"B":2,"D":2},
			{"I":2,"J":8,"B":2,"D":0},

			{"I":3,"J":0,"B":3,"D":0},
			{"I":3,"J":1,"B":3,"D":2},
			{"I":3,"J":2,"B":3,"D":0},
			{"I":3,"J":3,"B":4,"D":0},
			{"I":3,"J":4,"B":4,"D":1},
			{"I":3,"J":5,"B":4,"D":0},
			{"I":3,"J":6,"B":5,"D":0},
			{"I":3,"J":7,"B":5,"D":3},
			{"I":3,"J":8,"B":5,"D":0},

			{"I":4,"J":0,"B":3,"D":0},
			{"I":4,"J":1,"B":3,"D":0},
			{"I":4,"J":2,"B":3,"D":6},
			{"I":4,"J":3,"B":4,"D":5},
			{"I":4,"J":4,"B":4,"D":0},
			{"I":4,"J":5,"B":4,"D":4},
			{"I":4,"J":6,"B":5,"D":1},
			{"I":4,"J":7,"B":5,"D":0},
			{"I":4,"J":8,"B":5,"D":0},

			{"I":5,"J":0,"B":3,"D":0},
			{"I":5,"J":1,"B":3,"D":8},
			{"I":5,"J":2,"B":3,"D":0},
			{"I":5,"J":3,"B":4,"D":0},
			{"I":5,"J":4,"B":4,"D":7},
			{"I":5,"J":5,"B":4,"D":0},
			{"I":5,"J":6,"B":5,"D":0},
			{"I":5,"J":7,"B":5,"D":4},
			{"I":5,"J":8,"B":5,"D":0},

			{"I":6,"J":0,"B":6,"D":0},
			{"I":6,"J":1,"B":6,"D":7},
			{"I":6,"J":2,"B":6,"D":0},
			{"I":6,"J":3,"B":7,"D":0},
			{"I":6,"J":4,"B":7,"D":2},
			{"I":6,"J":5,"B":7,"D":0},
			{"I":6,"J":6,"B":8,"D":0},
			{"I":6,"J":7,"B":8,"D":6},
			{"I":6,"J":8,"B":8,"D":0},

			{"I":7,"J":0,"B":6,"D":0},
			{"I":7,"J":1,"B":6,"D":0},
			{"I":7,"J":2,"B":6,"D":1},
			{"I":7,"J":3,"B":7,"D":4},
			{"I":7,"J":4,"B":7,"D":0},
			{"I":7,"J":5,"B":7,"D":3},
			{"I":7,"J":6,"B":8,"D":5},
			{"I":7,"J":7,"B":8,"D":0},
			{"I":7,"J":8,"B":8,"D":0},

			{"I":8,"J":0,"B":6,"D":0},
			{"I":8,"J":1,"B":6,"D":0},
			{"I":8,"J":2,"B":6,"D":0},
			{"I":8,"J":3,"B":7,"D":0},
			{"I":8,"J":4,"B":7,"D":0},
			{"I":8,"J":5,"B":7,"D":0},
			{"I":8,"J":6,"B":8,"D":0},
			{"I":8,"J":7,"B":8,"D":0},
			{"I":8,"J":8,"B":8,"D":0}
	  ]
	}
`)

// "..3456789" +
// 	"456789123" +
// 	"789123456" +
// 	"..4365897" +
// 	"365897214" +
// 	"897214365" +
// 	"531642978" +
// 	"642978531" +
// 	"978531642" // 2 solutions puzzle
var terminalJson_9x9_2 = []byte(`
	{
		"E":9,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":0,"D":3},
			{"I":0,"J":3,"B":1,"D":4},
			{"I":0,"J":4,"B":1,"D":5},
			{"I":0,"J":5,"B":1,"D":6},
			{"I":0,"J":6,"B":2,"D":7},
			{"I":0,"J":7,"B":2,"D":8},
			{"I":0,"J":8,"B":2,"D":9},

			{"I":1,"J":0,"B":0,"D":4},
			{"I":1,"J":1,"B":0,"D":5},
			{"I":1,"J":2,"B":0,"D":6},
			{"I":1,"J":3,"B":1,"D":7},
			{"I":1,"J":4,"B":1,"D":8},
			{"I":1,"J":5,"B":1,"D":9},
			{"I":1,"J":6,"B":2,"D":1},
			{"I":1,"J":7,"B":2,"D":2},
			{"I":1,"J":8,"B":2,"D":3},

			{"I":2,"J":0,"B":0,"D":7},
			{"I":2,"J":1,"B":0,"D":8},
			{"I":2,"J":2,"B":0,"D":9},
			{"I":2,"J":3,"B":1,"D":1},
			{"I":2,"J":4,"B":1,"D":2},
			{"I":2,"J":5,"B":1,"D":3},
			{"I":2,"J":6,"B":2,"D":4},
			{"I":2,"J":7,"B":2,"D":5},
			{"I":2,"J":8,"B":2,"D":6},

			{"I":3,"J":0,"B":3,"D":0},
			{"I":3,"J":1,"B":3,"D":0},
			{"I":3,"J":2,"B":3,"D":4},
			{"I":3,"J":3,"B":4,"D":3},
			{"I":3,"J":4,"B":4,"D":6},
			{"I":3,"J":5,"B":4,"D":5},
			{"I":3,"J":6,"B":5,"D":8},
			{"I":3,"J":7,"B":5,"D":9},
			{"I":3,"J":8,"B":5,"D":7},

			{"I":4,"J":0,"B":3,"D":3},
			{"I":4,"J":1,"B":3,"D":6},
			{"I":4,"J":2,"B":3,"D":5},
			{"I":4,"J":3,"B":4,"D":8},
			{"I":4,"J":4,"B":4,"D":9},
			{"I":4,"J":5,"B":4,"D":7},
			{"I":4,"J":6,"B":5,"D":2},
			{"I":4,"J":7,"B":5,"D":1},
			{"I":4,"J":8,"B":5,"D":4},

			{"I":5,"J":0,"B":3,"D":8},
			{"I":5,"J":1,"B":3,"D":9},
			{"I":5,"J":2,"B":3,"D":7},
			{"I":5,"J":3,"B":4,"D":2},
			{"I":5,"J":4,"B":4,"D":1},
			{"I":5,"J":5,"B":4,"D":4},
			{"I":5,"J":6,"B":5,"D":3},
			{"I":5,"J":7,"B":5,"D":6},
			{"I":5,"J":8,"B":5,"D":5},

			{"I":6,"J":0,"B":6,"D":5},
			{"I":6,"J":1,"B":6,"D":3},
			{"I":6,"J":2,"B":6,"D":1},
			{"I":6,"J":3,"B":7,"D":6},
			{"I":6,"J":4,"B":7,"D":4},
			{"I":6,"J":5,"B":7,"D":2},
			{"I":6,"J":6,"B":8,"D":9},
			{"I":6,"J":7,"B":8,"D":7},
			{"I":6,"J":8,"B":8,"D":8},

			{"I":7,"J":0,"B":6,"D":6},
			{"I":7,"J":1,"B":6,"D":4},
			{"I":7,"J":2,"B":6,"D":2},
			{"I":7,"J":3,"B":7,"D":9},
			{"I":7,"J":4,"B":7,"D":7},
			{"I":7,"J":5,"B":7,"D":8},
			{"I":7,"J":6,"B":8,"D":5},
			{"I":7,"J":7,"B":8,"D":3},
			{"I":7,"J":8,"B":8,"D":1},

			{"I":8,"J":0,"B":6,"D":9},
			{"I":8,"J":1,"B":6,"D":7},
			{"I":8,"J":2,"B":6,"D":8},
			{"I":8,"J":3,"B":7,"D":5},
			{"I":8,"J":4,"B":7,"D":3},
			{"I":8,"J":5,"B":7,"D":1},
			{"I":8,"J":6,"B":8,"D":6},
			{"I":8,"J":7,"B":8,"D":4},
			{"I":8,"J":8,"B":8,"D":2}
	  ]
	}
`)

//
// "....7.94." +
// 	".7..9...5" +
// 	"3....5.7." +
// 	"..74..1.." +
// 	"463.8...." +
// 	".....7.8." +
// 	"8..7....." +
// 	"7......28" +
// 	".5..68..." // 188 solutions puzzle
var terminalJson_9x9_188 = []byte(`
	{
		"E":9,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":0,"D":0},
			{"I":0,"J":3,"B":1,"D":0},
			{"I":0,"J":4,"B":1,"D":7},
			{"I":0,"J":5,"B":1,"D":0},
			{"I":0,"J":6,"B":2,"D":9},
			{"I":0,"J":7,"B":2,"D":4},
			{"I":0,"J":8,"B":2,"D":0},

			{"I":1,"J":0,"B":0,"D":0},
			{"I":1,"J":1,"B":0,"D":7},
			{"I":1,"J":2,"B":0,"D":0},
			{"I":1,"J":3,"B":1,"D":0},
			{"I":1,"J":4,"B":1,"D":9},
			{"I":1,"J":5,"B":1,"D":0},
			{"I":1,"J":6,"B":2,"D":0},
			{"I":1,"J":7,"B":2,"D":0},
			{"I":1,"J":8,"B":2,"D":5},

			{"I":2,"J":0,"B":0,"D":3},
			{"I":2,"J":1,"B":0,"D":0},
			{"I":2,"J":2,"B":0,"D":0},
			{"I":2,"J":3,"B":1,"D":0},
			{"I":2,"J":4,"B":1,"D":0},
			{"I":2,"J":5,"B":1,"D":5},
			{"I":2,"J":6,"B":2,"D":0},
			{"I":2,"J":7,"B":2,"D":7},
			{"I":2,"J":8,"B":2,"D":0},

			{"I":3,"J":0,"B":3,"D":0},
			{"I":3,"J":1,"B":3,"D":0},
			{"I":3,"J":2,"B":3,"D":7},
			{"I":3,"J":3,"B":4,"D":4},
			{"I":3,"J":4,"B":4,"D":0},
			{"I":3,"J":5,"B":4,"D":0},
			{"I":3,"J":6,"B":5,"D":1},
			{"I":3,"J":7,"B":5,"D":0},
			{"I":3,"J":8,"B":5,"D":0},

			{"I":4,"J":0,"B":3,"D":4},
			{"I":4,"J":1,"B":3,"D":6},
			{"I":4,"J":2,"B":3,"D":3},
			{"I":4,"J":3,"B":4,"D":0},
			{"I":4,"J":4,"B":4,"D":8},
			{"I":4,"J":5,"B":4,"D":0},
			{"I":4,"J":6,"B":5,"D":0},
			{"I":4,"J":7,"B":5,"D":0},
			{"I":4,"J":8,"B":5,"D":0},

			{"I":5,"J":0,"B":3,"D":0},
			{"I":5,"J":1,"B":3,"D":0},
			{"I":5,"J":2,"B":3,"D":0},
			{"I":5,"J":3,"B":4,"D":0},
			{"I":5,"J":4,"B":4,"D":0},
			{"I":5,"J":5,"B":4,"D":7},
			{"I":5,"J":6,"B":5,"D":0},
			{"I":5,"J":7,"B":5,"D":8},
			{"I":5,"J":8,"B":5,"D":0},

			{"I":6,"J":0,"B":6,"D":8},
			{"I":6,"J":1,"B":6,"D":0},
			{"I":6,"J":2,"B":6,"D":0},
			{"I":6,"J":3,"B":7,"D":7},
			{"I":6,"J":4,"B":7,"D":0},
			{"I":6,"J":5,"B":7,"D":0},
			{"I":6,"J":6,"B":8,"D":0},
			{"I":6,"J":7,"B":8,"D":0},
			{"I":6,"J":8,"B":8,"D":0},

			{"I":7,"J":0,"B":6,"D":7},
			{"I":7,"J":1,"B":6,"D":0},
			{"I":7,"J":2,"B":6,"D":0},
			{"I":7,"J":3,"B":7,"D":0},
			{"I":7,"J":4,"B":7,"D":0},
			{"I":7,"J":5,"B":7,"D":0},
			{"I":7,"J":6,"B":8,"D":0},
			{"I":7,"J":7,"B":8,"D":2},
			{"I":7,"J":8,"B":8,"D":8},

			{"I":8,"J":0,"B":6,"D":0},
			{"I":8,"J":1,"B":6,"D":5},
			{"I":8,"J":2,"B":6,"D":0},
			{"I":8,"J":3,"B":7,"D":0},
			{"I":8,"J":4,"B":7,"D":6},
			{"I":8,"J":5,"B":7,"D":8},
			{"I":8,"J":6,"B":8,"D":0},
			{"I":8,"J":7,"B":8,"D":0},
			{"I":8,"J":8,"B":8,"D":0}
	  ]
	}
`)

// "." // 1 solutions puzzle
var terminalJson_1x1_1 = []byte(`
	{
		"E":1,
		"C":[
			{"I":0,"J":0,"B":0,"D":0}
	  ]
	}
`)

//
// "...." +
// 	".4.." +
// 	"2..." +
// 	"..43" // 0 solutions puzzle
var terminalJson_4x4_0 = []byte(`
	{
		"E":4,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":1,"D":0},
			{"I":0,"J":3,"B":1,"D":0},

			{"I":1,"J":0,"B":0,"D":0},
			{"I":1,"J":1,"B":0,"D":4},
			{"I":1,"J":2,"B":1,"D":0},
			{"I":1,"J":3,"B":1,"D":0},

			{"I":2,"J":0,"B":2,"D":2},
			{"I":2,"J":1,"B":2,"D":0},
			{"I":2,"J":2,"B":3,"D":0},
			{"I":2,"J":3,"B":3,"D":0},

			{"I":3,"J":0,"B":2,"D":0},
			{"I":3,"J":1,"B":2,"D":0},
			{"I":3,"J":2,"B":3,"D":4},
			{"I":3,"J":3,"B":3,"D":3}
	  ]
	}
`)

// "...." +
// 	".4.." +
// 	"2..." +
// 	"...3" // 3 solutions puzzle
var terminalJson_4x4_3 = []byte(`
	{
		"E":4,
		"C":[
			{"I":0,"J":0,"B":0,"D":0},
			{"I":0,"J":1,"B":0,"D":0},
			{"I":0,"J":2,"B":1,"D":0},
			{"I":0,"J":3,"B":1,"D":0},

			{"I":1,"J":0,"B":0,"D":0},
			{"I":1,"J":1,"B":0,"D":4},
			{"I":1,"J":2,"B":1,"D":0},
			{"I":1,"J":3,"B":1,"D":0},

			{"I":2,"J":0,"B":2,"D":2},
			{"I":2,"J":1,"B":2,"D":0},
			{"I":2,"J":2,"B":3,"D":0},
			{"I":2,"J":3,"B":3,"D":0},

			{"I":3,"J":0,"B":2,"D":0},
			{"I":3,"J":1,"B":2,"D":0},
			{"I":3,"J":2,"B":3,"D":0},
			{"I":3,"J":3,"B":3,"D":3}
	  ]
	}
`)
