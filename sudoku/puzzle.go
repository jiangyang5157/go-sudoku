package sudoku

import (
	"errors"
	"math/rand"
	"time"

	"github.com/jiangyang5157/go-dlx/dlx"
)

type puzzleData struct {
	dlx.X

	squares int // > 0
	edge    int // squares * squares
	cells   int // edge * edge

	offset1 int // 0
	offset2 int // offset1 + cells
	offset3 int // offset2 + cells
	offset4 int // offset3 + cells
}

/*
Constraints example: 9 x 9 puzzle (squares = 3)
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each squares must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
*/
func newPuzzleData(squares int) *puzzleData {
	edge := squares * squares
	cells := edge * edge
	offset1 := cells * 0
	offset2 := cells * 1
	offset3 := cells * 2
	offset4 := cells * 3
	return &puzzleData{
		squares: squares,
		edge:    edge,
		cells:   cells,

		offset1: offset1,
		offset2: offset2,
		offset3: offset3,
		offset4: offset4,
	}
}

func (p *puzzleData) build(digits []int) error {
	if digits == nil || len(digits) != p.cells {
		return errors.New("Invalid digits.")
	}

	p.X = *dlx.NewX(p.offset4 + p.cells)
	index := 0
	for row := 0; row < p.edge; row++ {
		for col := 0; col < p.edge; col++ {
			p.addDigit(digits[index], index, row, col, p.squareIndex(row, col))
			index++
		}
	}
	return nil
}

func (p *puzzleData) addDigit(digit int, index int, row int, col int, square int) {
	if digit >= 1 && digit <= p.edge {
		// valid digit
		p.AddRow([]int{
			p.offset1 + index + 1,
			p.offset2 + row*p.edge + digit,
			p.offset3 + col*p.edge + digit,
			p.offset4 + square*p.edge + digit})
	} else {
		// invalid digit, consider all possibilities
		for digit = 1; digit <= p.edge; digit++ {
			p.AddRow([]int{
				p.offset1 + index + 1,
				p.offset2 + row*p.edge + digit,
				p.offset3 + col*p.edge + digit,
				p.offset4 + square*p.edge + digit})
		}
	}
}

func (p *puzzleData) squareIndex(row int, col int) int {
	return row/p.squares*p.squares + col/p.squares
}

func (p *puzzleData) cellIndex(row int, col int) int {
	return row*p.edge + col
}

func (p *puzzleData) rowIndex(cell int) int {
	return cell / p.edge
}

func (p *puzzleData) colIndex(cell int) int {
	return cell % p.edge
}

func raw2digits(raw string) []int {
	digits := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		digits[i] = int(raw[i] - '0')
	}
	return digits
}

func digits2raw(digits []int) string {
	bs := make([]byte, len(digits))
	for i := 0; i < len(digits); i++ {
		bs[i] = byte(digits[i]) + '0'
	}
	return string(bs)
}

func digitsDisorder(digits []int) []int {
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(digits); i++ {
		random := rand.Intn(len(digits))
		digits[i], digits[random] = digits[random], digits[i]
	}
	return digits
}
