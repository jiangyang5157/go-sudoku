package sudoku

func SolveString(raw string) string {
	return string(SolveByte([]byte(raw)))
}

func SolveByte(raw []byte) []byte {
	t, _ := Raw2TerminalJson(raw)
	sol := SolveTerminalJson(t)
	ret, _ := TerminalJson2Raw(sol)
	return ret
}

func SolveTerminalJson(t *TerminalJson) *TerminalJson {
	var ret *TerminalJson
	newSudoku(t).initialize().solve(func(sol *TerminalJson) bool {
		ret = sol
		return true
	})
	return ret
}
