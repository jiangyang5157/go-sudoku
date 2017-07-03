package sudoku

import (
	"encoding/json"

	"github.com/jiangyang5157/go-dlx/dlx"
)

type sudoku struct {
	dlx.X

	puzzleJson PuzzleJson
}

func newSudoku(jsonRaw []byte) (*sudoku, error) {
	ret := &sudoku{}
	err := json.Unmarshal(jsonRaw, &ret.puzzleJson)
	if err != nil {
		return nil, err
	}
	err = ret.puzzleJson.validate()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *sudoku) initialize() {
	// TODO
}
