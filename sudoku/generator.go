package sudoku

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type GenMode int

const (
	REGULAR GenMode = iota
	IRREGULAR
)

type DIRECTION int

const (
	TOP DIRECTION = iota
	BOTTOM
	LEFT
	RIGHT
)

func GenString(edge int, mode GenMode, minSubGiven int, minTotalGiven int) string {
	return string(GenByte(edge, mode, minSubGiven, minTotalGiven))
}

func GenByte(edge int, mode GenMode, minSubGiven int, minTotalGiven int) []byte {
	t := genTerminal(edge, mode, minSubGiven, minTotalGiven)
	ret, _ := Terminal2Raw(t)
	return ret
}

func genTerminal(edge int, mode GenMode, minSubGiven int, minTotalGiven int) *Terminal {
	rand.Seed(time.Now().Unix())
	t := newTerminal(edge).genBlock(mode)
	if t == nil {
		return nil
	}
	t = t.genMaterial()
	t = solve(t)
	t = t.genPuzzle(minSubGiven, minTotalGiven)
	return t
}

func (t *Terminal) genBlock(mode GenMode) *Terminal {
	switch mode {
	case REGULAR:
		square := int(math.Sqrt(float64(t.E)))
		for i := 0; i < len(t.C); i++ {
			c := &t.C[i]
			c.B = (c.I/square)*square + c.J/square
		}
		return t
	case IRREGULAR:
		b := t.E - 1
		for i := 0; i < len(t.C); i++ {
			if b <= 0 {
				break
			}
			c := &t.C[i]
			if c.B == 0 {
				t = t.genCellBlockTop2BottomAndLeft2Right(i, b)
				b--
			}
		}
		return t
	default:
		return nil
	}
}

func (t *Terminal) genCellBlockTop2BottomAndLeft2Right(index int, block int) *Terminal {
	var c *Cell
	var trace []int
	count := 0

	c = &t.C[index]
	c.B = block
	trace = append(trace, index)
	count++

	for count < t.E {
		c = &t.C[trace[len(trace)-1]]
		dirs := []DIRECTION{}
		if c.I-1 >= 0 && t.Cell(c.I-1, c.J).B == 0 {
			dirs = append(dirs, TOP)
		} else if c.J-1 >= 0 && t.Cell(c.I, c.J-1).B == 0 {
			dirs = append(dirs, LEFT)
		} else {
			if c.I+1 < t.E && t.Cell(c.I+1, c.J).B == 0 {
				dirs = append(dirs, BOTTOM)
			}
			if c.J+1 < t.E && t.Cell(c.I, c.J+1).B == 0 {
				dirs = append(dirs, RIGHT)
			}
		}
		if len(dirs) == 0 {
			fmt.Printf("!!!!!!!! len(dirs) == 0, c.I=%v, c.J=%v, block=%v, trace=%v, c.B=%v\n", c.I, c.J, block, trace, c.B)
			if len(trace) <= 1 {
				return nil // TODO
			}
			trace = trace[:len(trace)-1]
			continue
		}

		fmt.Printf("#### trace=%v, c.index=%v\n", trace, t.Index(c.I, c.J))
		dir := dirs[rand.Intn(len(dirs))]
		switch dir {
		case TOP:
			c = t.Cell(c.I-1, c.J)
			c.B = block
			trace = append(trace, t.Index(c.I, c.J))
			count++
		case BOTTOM:
			c = t.Cell(c.I+1, c.J)
			c.B = block
			trace = append(trace, t.Index(c.I, c.J))
			count++
		case LEFT:
			c = t.Cell(c.I, c.J-1)
			c.B = block
			trace = append(trace, t.Index(c.I, c.J))
			count++
		case RIGHT:
			c = t.Cell(c.I, c.J+1)
			c.B = block
			trace = append(trace, t.Index(c.I, c.J))
			count++
		}
	}

	return t
}

// Fill diagonal square by random digits, returns the Terminal which should have solution
func (t *Terminal) genMaterial() *Terminal {
	tmp := make([]int, t.E)
	for i := range tmp {
		tmp[i] = i + 1 // [1, t.E]
	}
	square := int(math.Sqrt(float64(t.E)))
	for i := 0; i < t.E; i += square + 1 {
		digits := digitsDisorder(tmp)
		for j := 0; j < t.E; j++ {
			row := j/square + (i/square)*square
			col := j%square + (i/square)*square
			t.Cell(row, col).D = digits[j]
		}
	}
	return t
}

func (t *Terminal) genPuzzle(minSubGiven int, minTotalGiven int) *Terminal {
	remainTotalGiven := len(t.C)
	remainRowGiven := make([]int, t.E)
	remainColumnGiven := make([]int, t.E)

	tmp1 := make([]int, t.E)
	tmp2 := make([]int, t.E)
	for i := 0; i < t.E; i++ {
		remainRowGiven[i] = t.E
		remainColumnGiven[i] = t.E
		tmp1[i] = i // [0, t.E - 1]
		tmp2[i] = i // [0, t.E - 1]
	}

	s := newSudoku(t)
	dd1 := digitsDisorder(tmp1)
	for dd1i := 0; dd1i < t.E; dd1i++ {
		row := dd1[dd1i]
		dd2 := digitsDisorder(tmp2)
		for dd2i := 0; dd2i < t.E; dd2i++ {
			col := dd2[dd2i]
			switch {
			case remainTotalGiven <= minTotalGiven:
				continue
			case remainColumnGiven[col] <= minSubGiven:
				continue
			case remainRowGiven[row] <= minSubGiven:
				continue
			default:
				cell := t.Cell(row, col)
				cache := cell.D
				cell.D = 0 // remove the digit
				s.initialize()
				if s.hasUniqueSolution() {
					remainTotalGiven--
					remainColumnGiven[col]--
					remainRowGiven[row]--
				} else {
					cell.D = cache // resume the digit
				}
			}
		}
	}
	return t
}

func digitsDisorder(digits []int) []int {
	for i := 0; i < len(digits); i++ {
		random := rand.Intn(len(digits))
		digits[i], digits[random] = digits[random], digits[i]
	}
	return digits
}
