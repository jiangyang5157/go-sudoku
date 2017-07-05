package sudoku

import (
	"encoding/json"
	"fmt"
)

/*
E: Edge length of the TerminalJson
C: Each cell in the TerminalJson
I: Index i in the TerminalJson
J: Index j in the TerminalJson
B: Block that the Cell belongs to
D: Digit that the Cell hold
*/
type TerminalJson struct {
	fmt.Stringer
	E int    `json:"E"`
	C []Cell `json:"C"`
}

type Cell struct {
	I int `json:"I"`
	J int `json:"J"`
	B int `json:"B"`
	D int `json:"D"`
}

func NewTerminalJson(edge int) *TerminalJson {
	ret := &TerminalJson{E: edge}
	cells := edge * edge
	ret.C = make([]Cell, cells)
	index := 0
	for i := 0; i < edge; i++ {
		for j := 0; j < edge; j++ {
			ret.C[index].I = i
			ret.C[index].J = j
			index++
		}
	}
	return ret
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

func (t *TerminalJson) Cell(i int, j int) *Cell {
	return &t.C[t.Index(i, j)]
}

func (t *TerminalJson) Index(i int, j int) int {
	return i*t.E + j
}

func (t *TerminalJson) String() string {
	const ASCII_0 = '0'
	ret := fmt.Sprintf("Terminal: E=%c, []{D(I,J,B)}=\n", t.E+ASCII_0)
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			c := t.Cell(i, j)
			ret += fmt.Sprintf("%c(%c,%c,%c),", c.D+ASCII_0, c.I+ASCII_0, c.J+ASCII_0, c.B+ASCII_0)
		}
		ret += "\n"
	}
	return ret
}
