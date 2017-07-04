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
func (s *sudoku) initialize() {
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
}