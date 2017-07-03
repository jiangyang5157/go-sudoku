package sudoku

import "testing"

func Test_newSudoku(t *testing.T) {
	s, err := newSudoku(jsonRaw)
	if err != nil {
		t.Error(err)
	}
	if s.puzzleJson.Config.Edge != 2 {
		t.Error("Config.Edge should be 2")
	}
	if len(s.puzzleJson.Terminal) != 2 {
		t.Error("len(s.puzzleJson.Terminal) should be 2")
	}
	if len(s.puzzleJson.Terminal[1].Row) != 2 {
		t.Error("len(s.puzzleJson.Terminal[1].Row) should be 2")
	}
	if len(s.puzzleJson.Terminal[1].Row[1].Column) != 2 {
		t.Error("len(s.puzzleJson.Terminal[1].Row[1].Column) should be 1")
	}
}
