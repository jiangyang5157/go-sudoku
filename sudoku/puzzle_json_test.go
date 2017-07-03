package sudoku

import (
	"encoding/json"
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

func Test_PuzzleJson_from_jsonData(t *testing.T) {
	var puzzleJson PuzzleJson
	err := json.Unmarshal(jsonRaw, &puzzleJson)
	if err != nil {
		t.Error(err)
	}
	err = puzzleJson.validate()
	if err != nil {
		t.Error(err)
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
	if len(puzzleJson.Terminal[1].Row[1].Column) != 2 {
		t.Error("len(puzzleJson.Terminal[1].Row[1].Column) should be 1")
	}
}

func Test_PuzzleJson_to_jsonData(t *testing.T) {
	var puzzleJson PuzzleJson
	json.Unmarshal(jsonRaw, &puzzleJson)
	jsonData2, _ := json.Marshal(puzzleJson)
	var puzzleJson2 PuzzleJson
	json.Unmarshal(jsonData2, &puzzleJson2)

	if !reflect.DeepEqual(puzzleJson2, puzzleJson) {
		t.Error("reflect.DeepEqual(puzzleJson2, puzzleJson) should be true")
	}
}
