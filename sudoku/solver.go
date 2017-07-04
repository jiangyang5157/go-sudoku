package sudoku

func SolveString(raw string) string {
	return string(SolveByte([]byte(raw)))
}

func SolveByte(raw []byte) []byte {
	t, _ := Raw2Terminal(raw)
	sol := solve(t)
	ret, _ := Terminal2Raw(sol)
	return ret
}

func solve(t *Terminal) *Terminal {
	var ret *Terminal
	newSudoku(t).initialize().solve(func(sol *Terminal) bool {
		ret = sol
		return true
	})
	return ret
}
