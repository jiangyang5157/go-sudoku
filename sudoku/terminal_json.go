package sudoku

import (
	"encoding/json"
	"fmt"
)

/*
E: Edge length of the TerminalJson
C: Each cell in the TerminalJson. They are following left-to-right(columns) && up-to-down(rows) order.
B: Block that the Cell belongs to
D: Digit that the Cell hold
*/
type TerminalJson struct {
	fmt.Stringer
	E int    `json:"E"`
	C []Cell `json:"C"`
}

type Cell struct {
	B int `json:"B"`
	D int `json:"D"`
}

func NewTerminalJson(edge int) *TerminalJson {
	return &TerminalJson{
		E: edge,
		C: make([]Cell, edge*edge),
	}
}

func Raw2TerminalJson(raw []byte) (*TerminalJson, error) {
	var ret *TerminalJson
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func TerminalJson2Raw(t *TerminalJson) ([]byte, error) {
	ret, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TerminalJson) Clone() *TerminalJson {
	raw, _ := TerminalJson2Raw(t)
	ret, _ := Raw2TerminalJson(raw)
	return ret
}

func (t *TerminalJson) Cell(row int, col int) *Cell {
	return &t.C[t.Index(row, col)]
}

func (t *TerminalJson) Index(row int, col int) int {
	return row*t.E + col
}

func (t *TerminalJson) RowCol(index int) (row, col int) {
	row = index / t.E
	col = index % t.E
	return
}

func (t *TerminalJson) String() string {
	const ASCII_0 = '0'
	ret := fmt.Sprintf("TerminalJson: E=%c, {D[B]...}=\n", t.E+ASCII_0)
	index := 0
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			c := t.C[index]
			index++
			ret += fmt.Sprintf("%c[%c],", c.D+ASCII_0, c.B+ASCII_0)
		}
		ret += "\n"
	}
	return ret
}
