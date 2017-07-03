package sudoku

import (
	"encoding/json"
	"reflect"
	"testing"
)

var jsonData = []byte(`
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
              }
            ]
          },
          {
            "ColumnJson":[
              {
                "Digit":2,
                "Block":2
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
              }
            ]
          },
          {
            "ColumnJson":[
              {
                "Digit":1,
                "Block":2
              }
            ]
          }
        ]
      }
    ]
  }
`)

func Test_PuzzleJson_from_jsonData(t *testing.T) {
	var puzzleJson PuzzleJson
	err := json.Unmarshal(jsonData, &puzzleJson)
	if err != nil {
		t.Error("Invalid json data:\n%v", jsonData)
	}
	if puzzleJson.Config.Edge != 2 {
		t.Error("Config.Edge should be 2")
	}
	if len(puzzleJson.Terminal) != 2 {
		t.Error("len(puzzleJson.Terminal) should be 2")
	}
	if len(puzzleJson.Terminal[1].Row) != 2 {
		t.Error("len(puzzleJson.Terminal[1].Row) should be 2")
	}
	if len(puzzleJson.Terminal[1].Row[1].Column) != 1 {
		t.Error("len(puzzleJson.Terminal[1].Row[1].Column) should be 1")
	}
	if puzzleJson.Terminal[1].Row[1].Column[0].Digit != 1 {
		t.Error("Terminal[1].Row[1].Column[0].Digit should be 1")
	}
	if puzzleJson.Terminal[1].Row[1].Column[0].Block != 2 {
		t.Error("Terminal[1].Row[1].Column[0].Digit should be 2")
	}
}

func Test_PuzzleJson_to_jsonData(t *testing.T) {
	var puzzleJson PuzzleJson
	json.Unmarshal(jsonData, &puzzleJson)
	jsonData2, _ := json.Marshal(puzzleJson)
	var puzzleJson2 PuzzleJson
	json.Unmarshal(jsonData2, &puzzleJson2)

	if !reflect.DeepEqual(puzzleJson2, puzzleJson) {
		t.Error("reflect.DeepEqual(puzzleJson2, puzzleJson) should be true")
	}
}
