package sudoku

import (
	"encoding/json"
	"fmt"
)

/*
E: Edge length of the Terminal
C: Each cell in the Terminal
I: Index i in the Terminal
J: Index j in the Terminal
B: Block that the Cell belongs to
D: Digit that the Cell hold
*/
type Terminal struct {
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

func Raw2Terminal(raw []byte) (*Terminal, error) {
	var ret *Terminal
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func Terminal2Raw(t *Terminal) ([]byte, error) {
	ret, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *Terminal) Clone() *Terminal {
	raw, _ := Terminal2Raw(t)
	ret, _ := Raw2Terminal(raw)
	return ret
}

func (t *Terminal) Cell(i int, j int) *Cell {
	return &t.C[i*t.E+j]
}

func (t *Terminal) String() string {
	var ret string
	const ASCII_0 = '0'
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			c := t.Cell(i, j)
			ret += fmt.Sprintf("%c(%c),", c.D+ASCII_0, c.B+ASCII_0)
		}
		ret += "\n"
	}
	return ret
}
