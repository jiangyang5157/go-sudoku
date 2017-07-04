package sudoku

func SolveString(raw string) string {
	return string(SolveByte([]byte(raw)))
}

func SolveByte(raw []byte) []byte {
	rawT, _ := Raw2Terminal(raw)
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
