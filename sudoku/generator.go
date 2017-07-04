package sudoku

import (
	"math/rand"
	"time"
)

type Gen_Mode int

const (
	NORMAL Gen_Mode = iota // 0
	RANDOM                 // 1
)

func GenString(edge int, mode Gen_Mode, minSubGiven int, minTotalGiven int) string {
	return string(GenByte(edge, mode, minSubGiven, minTotalGiven))
}

func GenByte(edge int, mode Gen_Mode, minSubGiven int, minTotalGiven int) []byte {
	t := genTerminal(edge, mode, minSubGiven, minTotalGiven)
	ret, _ := Terminal2Raw(t)
	return ret
}

func genTerminal(edge int, mode Gen_Mode, minSubGiven int, minTotalGiven int) *Terminal {
	t := genMaterial(edge).genBlock(mode)
	t = solve(t)
	if t == nil {
		return nil
	}
	t = t.genPuzzle(minSubGiven, minTotalGiven)
	return t
}

func genMaterial(edge int) *Terminal {
	t := newTerminal(edge)
	// TODO
	return t
}

func (t *Terminal) genBlock(mode Gen_Mode) *Terminal {
	// TODO
	return t
}

func (t *Terminal) genPuzzle(minSubGiven int, minTotalGiven int) *Terminal {
	// TODO
	return t
}

func digitsDisorder(digits []int) []int {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(digits); i++ {
		random := rand.Intn(len(digits))
		digits[i], digits[random] = digits[random], digits[i]
	}
	return digits
}

// func (p *puzzle) generateTerminalPuzzle() []int {
// 	ret := make([]int, p.cells)
// 	digits := make([]int, p.cells)
//
// 	tmp := make([]int, p.edge)
// 	for i := range tmp {
// 		tmp[i] = i + 1 // 1,2,...,p.edge
// 	}
//
// 	// rare zero solution happens, particularly for 2x2 puzzle
// 	ok := false
// 	for {
// 		if ok {
// 			break
// 		}
//
// 		// fill diagonal squares by random digits
// 		for i := 0; i < p.edge; i += p.squares + 1 {
// 			randomDigits := digitsDisorder(tmp)
// 			for j := 0; j < p.edge; j++ {
// 				row := j/p.squares + (i/p.squares)*p.squares
// 				col := j%p.squares + (i/p.squares)*p.squares
// 				digits[p.cellIndex(row, col)] = randomDigits[j]
// 			}
// 		}
//
// 		p.build(digits)
// 		p.Search(func(sol dlx.Solution) bool {
// 			for _, nd := range sol {
// 				nd_row_col_index := nd.Row.Col.Index             // [offset1 + 1, offset2]
// 				nd_row_right_col_index := nd.Row.Right.Col.Index // [offset2 + 1, offset3]
// 				index := nd_row_col_index - 1
// 				digit := (nd_row_right_col_index - 1) % p.edge // [0, cells - 1]
// 				ret[index] = digit + 1
// 			}
// 			ok = true
// 			return true
// 		})
// 	}
//
// 	return ret
// }
//
// // Generate unique solution puzzle.
// func GeneratePuzzle(squares int, minSubGiven int, minTotalGiven int) string {
// 	p := newPuzzle(squares)
//
// 	remainTotalGiven := p.cells
// 	remainRowGiven := make([]int, p.edge)
// 	remainColumnGiven := make([]int, p.edge)
//
// 	tmp1 := make([]int, p.edge)
// 	tmp2 := make([]int, p.edge)
//
// 	for i := 0; i < p.edge; i++ {
// 		remainRowGiven[i] = p.edge
// 		remainColumnGiven[i] = p.edge
// 		tmp1[i] = i // 0,1,...,p.edge - 1
// 		tmp2[i] = i
// 	}
//
// 	terminalPuzzle := p.generateTerminalPuzzle()
// 	dd1 := digitsDisorder(tmp1)
// 	for dd1i := 0; dd1i < p.edge; dd1i++ {
// 		row := dd1[dd1i]
// 		dd2 := digitsDisorder(tmp2)
// 		for dd2i := 0; dd2i < p.edge; dd2i++ {
// 			col := dd2[dd2i]
// 			switch {
// 			case remainTotalGiven <= minTotalGiven:
// 				continue
// 			case remainColumnGiven[col] <= minSubGiven:
// 				continue
// 			case remainRowGiven[row] <= minSubGiven:
// 				continue
// 			default:
// 				cell := p.cellIndex(row, col)
// 				digit := terminalPuzzle[cell]
// 				// destroy the digit
// 				terminalPuzzle[cell] = 0
// 				p.build(terminalPuzzle)
// 				if p.HasUniqueSolution() {
// 					remainTotalGiven--
// 					remainColumnGiven[col]--
// 					remainRowGiven[row]--
// 				} else {
// 					// resume the digit
// 					terminalPuzzle[cell] = digit
// 				}
// 			}
// 		}
// 	}
//
// 	return digits2raw(terminalPuzzle)
// }
