package sudoku

import (
	"reflect"
	"testing"
)

var jsonRaw = []byte(`
  {
    "ConfigJson" : {
      "Edge":2
    },
    "TerminalJson": [
      {
        "RowJson":[
          {
            "ColumnJson":[
              {
                "Digit":1,
                "Block":1
              },
							{
                "Digit":2,
                "Block":1
              }
            ]
          },
          {
            "ColumnJson":[
              {
                "Digit":3,
                "Block":1
              },
							{
                "Digit":4,
                "Block":1
              }
            ]
          }
        ]
      },
			{
        "RowJson":[
          {
            "ColumnJson":[
              {
                "Digit":2,
                "Block":1
              },
							{
                "Digit":1,
                "Block":1
              }
            ]
          },
          {
            "ColumnJson":[
              {
                "Digit":4,
                "Block":1
              },
							{
                "Digit":3,
                "Block":1
              }
            ]
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
	err = puzzle.Validate()
	if err != nil {
		t.Error(err)
	}
	if puzzle.Config.Edge != 2 {
		t.Error("Config.Edge should be 2")
	}
	if len(puzzle.Terminal) != 2 {
		t.Error("len(puzzleJson.Terminal) should be 2")
	}
	if len(puzzle.Terminal[1].Row) != 2 {
		t.Error("len(puzzleJson.Terminal[1].Row) should be 2")
	}
	if len(puzzle.Terminal[1].Row[1].Column) != 2 {
		t.Error("len(puzzleJson.Terminal[1].Row[1].Column) should be 1")
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
