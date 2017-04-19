package sudoku

import (
	"fmt"
	"testing"
)

func Test_Ascii(t *testing.T) {
	fmt.Printf("%v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v\n",
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, ".")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n",
		int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
}

func Test_digits2raw2digits(t *testing.T) {
	var digits []int = []int{1, 2, 13, 14, 5, 11, 7, 10, 2}
	fmt.Println("digits", digits)
	raw := digits2raw(digits)
	fmt.Println("digits -> raw", raw)
	fmt.Println("raw -> digits", raw2digits(raw))
}

func Test_digitsDisorder(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	digitsDisorder(digits)
	fmt.Printf("digitsDisorder %d \n", digits)
}

func printPuzzleDigits(squares int, digits []int) {
	printPuzzleRaw(squares, digits2raw(digits))
}

func printPuzzleRaw(squares int, raw string) {
	switch squares {
	case 1:
		fmt.Printf("%c\n", raw[0])
	case 2:
		for r, i := 0, 0; r < 4; r, i = r+1, i+4 {
			fmt.Printf("%c %c | %c %c\n",
				raw[i], raw[i+1],
				raw[i+2], raw[i+3])
			if r == 1 {
				fmt.Println("----+----")
			}
		}
	case 3:
		for r, i := 0, 0; r < 9; r, i = r+1, i+9 {
			fmt.Printf("%c %c %c | %c %c %c | %c %c %c\n",
				raw[i], raw[i+1], raw[i+2],
				raw[i+3], raw[i+4], raw[i+5],
				raw[i+6], raw[i+7], raw[i+8])
			if r == 2 || r == 5 {
				fmt.Println("------+-------+------")
			}
		}
	case 4:
		// Ascii for chars after '9' can convert into integer by minus '0' for clear display.
		for r, i := 0, 0; r < 16; r, i = r+1, i+16 {
			fmt.Printf("%c %c %c %c | %c %c %c %c | %c %c %c %c | %c %c %c %c\n",
				raw[i], raw[i+1], raw[i+2], raw[i+3],
				raw[i+4], raw[i+5], raw[i+6], raw[i+7],
				raw[i+8], raw[i+9], raw[i+10], raw[i+11],
				raw[i+12], raw[i+13], raw[i+14], raw[i+15])
			if r == 3 || r == 7 || r == 11 {
				fmt.Println("--------+---------+---------+--------")
			}
		}
	default:
		fmt.Printf("squares: %v\nraw: %v\n", squares, raw)
	}
	fmt.Println("================================================================")
}
