package sudoku

import (
	"reflect"
	"testing"
)

var jsonRaw = []byte(`
{
	"E":2,
  "T": [
    {
      "C":[
        {
          "I":0,
          "J":0,
          "B":0,
          "D":1
        },
        {
          "I":0,
          "J":1,
          "B":0,
          "D":2
        },
        {
          "I":1,
          "J":0,
          "B":1,
          "D":1
        },
        {
          "I":1,
          "J":1,
          "B":1,
          "D":0
        }
      ]
    }
  ]
}
`)

func Test_Raw2Puzzle(t *testing.T) {
	puzzle, err := Raw2Puzzle(jsonRaw)
	if err != nil {
		t.Error(err)
	}
	if !puzzle.Validate() || puzzle.E != 2 || puzzle.T[0].C[3].I != 1 || puzzle.T[0].C[3].J != 1 || puzzle.T[0].C[3].B != 1 || puzzle.T[0].C[3].D != 0 {
		t.Error("Incorrect puzzle data")
	}
}

func Test_Puzzle2Raw(t *testing.T) {
	puzzle, _ := Raw2Puzzle(jsonRaw)
	raw2, _ := Puzzle2Raw(puzzle)
	puzzle2, _ := Raw2Puzzle(raw2)
	if !reflect.DeepEqual(puzzle2, puzzle) {
		t.Error("reflect.DeepEqual(puzzle2, puzzle) should be true")
	}
}
