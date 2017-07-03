package sudoku

import "encoding/json"

type PuzzleJson struct {
	T []struct {
		E int `json:"E"` // Edge length
		C []struct {
			I int `json:"I"` // Index i in the Terminal
			J int `json:"J"` // Index j in the Terminal
			B int `json:"B"` // Block that the Cell belongs to
			D int `json:"D"` // Digit that the Cell hold
		} `json:"C"` // Cell
	} `json:"T"` // Terminal
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
