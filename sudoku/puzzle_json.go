package sudoku

import (
	"encoding/json"
	"errors"
)

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

func Raw2Puzzle(raw []byte) (*PuzzleJson, error) {
	var ret *PuzzleJson
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func Puzzle2Raw(puzzle *PuzzleJson) ([]byte, error) {
	ret, err := json.Marshal(puzzle)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (p *PuzzleJson) Validate() error {
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
