package sudoku

import "github.com/jiangyang5157/go-dlx/dlx"

type sudoku struct {
	x dlx.X
	t *Terminal
}

func newSudoku(terminal *Terminal) *sudoku {
	return &sudoku{t: terminal}
}

/*
Constraints example: 9 x 9 puzzle
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each block must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
*/
func (s *sudoku) initialize() *sudoku {
	cells := s.t.E * s.t.E
	cellConstraintOffset := 0
	rowConstraintOffset := cellConstraintOffset + cells
	columnConstraintOffset := rowConstraintOffset + cells
	blockConstraintOffset := columnConstraintOffset + cells

	s.x = *dlx.NewX(blockConstraintOffset + cells)
	for i := 0; i < cells; i++ {
		c := s.t.C[i]
		if c.D >= 1 && c.D <= s.t.E {
			// valid digit
			s.x.AddRow([]int{
				cellConstraintOffset + i + 1,
				rowConstraintOffset + c.I*s.t.E + c.D,
				columnConstraintOffset + c.J*s.t.E + c.D,
				blockConstraintOffset + c.B*s.t.E + c.D,
			})
		} else {
			// invalid digit, consider all the possible digits
			for n := 1; n <= s.t.E; n++ {
				s.x.AddRow([]int{
					cellConstraintOffset + i + 1,
					rowConstraintOffset + c.I*s.t.E + n,
					columnConstraintOffset + c.J*s.t.E + n,
					blockConstraintOffset + c.B*s.t.E + n,
				})
			}
		}
	}
	return s
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

func (s *sudoku) hasUniqueSolution() bool {
	found := 0
	s.x.Search(func(dlx.Solution) bool {
		found++
		return found > 1
	})
	return found == 1
}
