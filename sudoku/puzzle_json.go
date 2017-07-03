package sudoku

type PuzzleJson struct {
	Config   ConfigJson     `json:"ConfigJson"`
	Terminal []TerminalJson `json:"TerminalJson"`
}

type ConfigJson struct {
	Edge int `json:"Edge"`
}

type TerminalJson struct {
	Row []RowJson `json:"RowJson"`
}

type RowJson struct {
	Column []ColumnJson `json:"ColumnJson"`
}

type ColumnJson struct {
	Digit int `json:"Digit"`
	Block int `json:"Block"`
}
