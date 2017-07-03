package sudoku

import "encoding/json"

/*
E: Edge length of the Terminal
C: Each cell in the Terminal
I: Index i in the Terminal
J: Index j in the Terminal
B: Block that the Cell belongs to
D: Digit that the Cell hold
*/
type Terminal struct {
	E int `json:"E"`
	C []struct {
		I int `json:"I"`
		J int `json:"J"`
		B int `json:"B"`
		D int `json:"D"`
	} `json:"C"`
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
