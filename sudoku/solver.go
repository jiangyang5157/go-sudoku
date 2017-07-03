package sudoku

import "github.com/jiangyang5157/go-dlx/dlx"

// prefix < '0' && prefix != whatever representing unknown digit in the raw
const SOLUTION_PREFIX byte = '#'

func SolveRaw(squares int, raw string, solutions int) string {
	return SolveDigits(squares, raw2digits(raw), solutions)
}

func SolveDigits(squares int, digits []int, solutions int) string {
	p := newPuzzle(squares)
	err := p.build(digits)
	if err != nil {
		return ""
	}
	return p.solve(solutions)
}

func (p *puzzleData) solve(solutions int) string {
	var ret []byte
	count := 0
	p.Search(func(sol dlx.Solution) bool {
		bs := make([]byte, p.cells)
		for _, nd := range sol {
			nd_row_col_index := nd.Row.Col.Index             // [offset1 + 1, offset2]
			nd_row_right_col_index := nd.Row.Right.Col.Index // [offset2 + 1, offset3]
			index := nd_row_col_index - 1
			digit := (nd_row_right_col_index - 1) % p.edge // [0, cells - 1]
			bs[index] = byte(digit) + '1'
		}
		ret = append(ret, SOLUTION_PREFIX)
		ret = append(ret, bs...)
		count++
		return count >= solutions
	})
	return string(ret)
}

func (p *puzzleData) HasUniqueSolution() bool {
	count := 0
	p.Search(func(sol dlx.Solution) bool {
		count++
		return count > 1
	})
	return count == 1
}
