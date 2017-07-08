package sudoku

import (
	"math"
	"math/rand"
	"time"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
	"github.com/jiangyang5157/golang-start/data/stack"
)

type GeneratorMode int

const (
	SQUARE GeneratorMode = iota
	RANDOM
	IRREGULAR
)

func GenString(edge int, mode GeneratorMode, minSubGiven int, minTotalGiven int) string {
	return string(GenByte(edge, mode, minSubGiven, minTotalGiven))
}

func GenByte(edge int, mode GeneratorMode, minSubGiven int, minTotalGiven int) []byte {
	t := genTerminal(edge, mode, minSubGiven, minTotalGiven)
	ret, _ := TerminalJson2Raw(t)
	return ret
}

func genTerminal(edge int, mode GeneratorMode, minSubGiven int, minTotalGiven int) *TerminalJson {
	t := NewTerminalJson(edge).genBlock(mode)
	t = t.genMaterial()
	t = solve(t)
	t = t.genPuzzle(minSubGiven, minTotalGiven)
	return t
}

func (t *TerminalJson) genBlock(mode GeneratorMode) *TerminalJson {
	switch mode {
	case SQUARE:
		square := int(math.Sqrt(float64(t.E)))
		for i := 0; i < len(t.C); i++ {
			c := &t.C[i]
			row, col := t.Row(i), t.Col(i)
			c.B = (row/square)*square + col/square
		}
		return t
	case RANDOM:
		rand.Seed(time.Now().Unix())
		tmp := disorderDigits(increasingDigits(0, len(t.C)-1))
		b := t.E - 1
		index := 0
		for i := 0; i < t.E; i++ {
			if b < 1 {
				break
			}
			for j := 0; j < t.E; j++ {
				t.C[tmp[index]].B = b
				index++
			}
			b--
		}
		return t
	case IRREGULAR:
		// TODO temp alg
		g := newGraph(t)
		b := t.E - 1
		tempMaxRun := 0
		for b > 0 && tempMaxRun <= 2 {
			tempMaxRun++
			// up-to-down, left-to-right
			for i := 0; i < len(t.C); i++ {
				if b == 0 {
					// rest of cells belongs to block 0
					break
				}
				if t.C[i].B > 0 {
					// already be assigned to a particular block
					continue
				}
				if genIrregularBlock(t, g, b, i) {
					b--
				}
			}
		}
		if b > 0 {
			count := 0
			for i := 0; i < len(t.C); i++ {
				if b == 0 {
					break
				}
				t.C[i].B = b
				count++
				if count == t.E {
					b--
				}
			}
		}
		return t
	default:
		return nil
	}
}

func genIrregularBlock(t *TerminalJson, g graph.Graph, block int, begin int) bool {
	remain := len(t.C) - (t.E-1-block)*t.E
	tgtRemain := remain - t.E

	// Because of the up-to-down and left-to-right order,
	// C[begin] is valid to be the first one of target block.
	trace := stack.NewStack()
	nbsOfBegin := srcNeighbours(t, g, begin)
	for _, nb := range nbsOfBegin {
		unlink(g, index2id(nb), index2id(begin))
	}
	trace.Push(begin)
	t.C[begin].B = block
	remain--

	var undo []int
	rand.Seed(time.Now().Unix())
	for remain > tgtRemain {
		if trace.IsEmpty() {
			break
		}
		ok := false
		for !ok {
			if trace.IsEmpty() {
				break
			}

			index := trace.Peek().(int)
			nbs := disorderDigits(tgtNeighbours(t, g, index))

			for _, nb := range nbs {
				trace.Push(nb)
				nbsOfNbs := srcNeighbours(t, g, nb)
				for _, nbOfNbs := range nbsOfNbs {
					unlink(g, index2id(nbOfNbs), index2id(nb))
				}
				t.C[nb].B = block

				// as the begin of dfs
				var indexOfAnyBlock0 int
				for i := len(t.C) - 1; i >= 0; i-- {
					if t.C[i].B == 0 {
						indexOfAnyBlock0 = i
						break
					}
				}
				visited := reachableNum(g, index2id(indexOfAnyBlock0))

				if visited == remain-1 {
					ok = true
					break
				} else {
					for _, nbOfNbs := range nbsOfNbs {
						link(g, index2id(nbOfNbs), index2id(nb))
					}
					t.C[nb].B = 0
					trace.Pop()
				}
			}
			if !ok {
				undo = append(undo, trace.Pop().(int))
			}
		}
		remain--
	}

	ret := !trace.IsEmpty()
	if !ret {
		for _, v := range undo {
			nbs := t.Neighbours(v)
			for _, v2 := range nbs {
				if t.C[v2].B == block || t.C[v2].B == 0 {
					link(g, index2id(v), index2id(v2))
					link(g, index2id(v2), index2id(v))
				}
			}
		}

		for _, nb := range nbsOfBegin {
			link(g, index2id(nb), index2id(begin))
		}
		t.C[begin].B = 0
	}

	return ret
}

func reachableNum(g graph.Graph, id graph.Id) int {
	visited := 0
	traversal.Dfs(g, id, func(nd graph.Node) bool {
		visited++
		return false
	})
	return visited
}

// Fill diagonal square by random digits, returns the Terminal which should have solution
func (t *TerminalJson) genMaterial() *TerminalJson {
	tmp := increasingDigits(1, t.E)
	square := int(math.Sqrt(float64(t.E)))
	for i := 0; i < t.E; i += square + 1 {
		digits := disorderDigits(tmp)
		for j := 0; j < t.E; j++ {
			row := j/square + (i/square)*square
			col := j%square + (i/square)*square
			t.Cell(row, col).D = digits[j]
		}
	}
	return t
}

func (t *TerminalJson) genPuzzle(minSubGiven int, minTotalGiven int) *TerminalJson {
	remainTotalGiven := len(t.C)
	remainRowGiven := make([]int, t.E)
	remainColumnGiven := make([]int, t.E)
	tmp1 := make([]int, t.E)
	tmp2 := make([]int, t.E)
	for i := 0; i < t.E; i++ {
		remainRowGiven[i] = t.E
		remainColumnGiven[i] = t.E
		tmp1[i] = i // [0, t.E - 1]
		tmp2[i] = i
	}

	s := newSudoku(t)
	rand.Seed(time.Now().Unix())
	dd1 := disorderDigits(tmp1)

	for dd1i := 0; dd1i < t.E; dd1i++ {
		row := dd1[dd1i]
		dd2 := disorderDigits(tmp2)
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

func increasingDigits(start, end int) []int {
	ret := make([]int, end-start+1)
	for i := range ret {
		ret[i] = i + start
	}
	return ret
}

func disorderDigits(src []int) []int {
	for i := 0; i < len(src); i++ {
		random := rand.Intn(len(src))
		src[i], src[random] = src[random], src[i]
	}
	return src
}
