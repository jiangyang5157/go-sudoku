package sudoku

import "errors"

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

func (p *PuzzleJson) validate() error {
	for _, t := range p.Terminal {
		if len(t.Row) != p.Config.Edge {
			return errors.New("Invalid row size.")
		}
		for _, r := range t.Row {
			if len(r.Column) != p.Config.Edge {
				return errors.New("Invalid column size.")
			}
		}
	}
	return nil
}
