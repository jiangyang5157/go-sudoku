package sudoku

import "github.com/jiangyang5157/go-dlx/dlx"

func Solve(rawJson []byte) []byte {
	rawT, _ := Raw2Terminal(rawJson)
	s := newSudoku(rawT)
	s.initialize()
	var retT *Terminal
	s.solve(func(t *Terminal) bool {
		retT = t
		return true
	})
	retJson, _ := Terminal2Raw(retT)
	return retJson
}

func (s *sudoku) solve(f func(*Terminal) bool) {
	s.x.Search(func(sol dlx.Solution) bool {
		t := s.t.Clone()
		for _, nd := range sol {
			nd_row_col_index := nd.Row.Col.Index             // [cellConstraintOffset + 1, rowConstraintOffset]
			nd_row_right_col_index := nd.Row.Right.Col.Index // [rowConstraintOffset + 1, columnConstraintOffset]
			index := nd_row_col_index - 1                    // [0, cells - 1]
			digit := (nd_row_right_col_index-1)%s.t.E + 1    // [0, edge - 1]
			t.C[index].D = digit
		}
		return f(t)
	})
}
