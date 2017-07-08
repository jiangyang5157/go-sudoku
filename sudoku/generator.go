package sudoku

import (
	"math"
	"math/rand"
	"time"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
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
		// g := newGraph(t)
		// tgtBlock := t.E - 1
		// // gen block follow up-to-down and left-to-right order
		// trace := stack.NewStack()
		// for i := 0; i < len(t.C); i++ {
		// 	if tgtBlock <= 0 {
		// 		// rest of cells belongs to block 0
		// 		break
		// 	}
		// 	if t.C[i].B > 0 {
		// 		// already be assigned to a particular block
		// 		continue
		// 	}
		// 	genIrregularBlock(t, g, i, tgtBlock, trace)
		// 	fmt.Printf("trace: %v\n", trace)
		// 	tgtBlock--
		// }
		//----
		// rand.Seed(time.Now().Unix())
		// tmp := disorderDigits(increasingDigits(0, len(t.C)-1))
		// todo := make(map[int]bool, len(tmp))
		// for _, index := range tmp {
		// 	todo[index] = true
		// }
		// g := newGraph(t)
		// b := t.E - 1
		// for b > 0 {
		// 	for k, v := range todo {
		// 		if !v {
		// 			continue
		// 		}
		//
		// 		result := markIrregularBlock(t, g, k, b)
		// 		if len(result) != t.E {
		// 			continue
		// 		}
		//
		// 		for ret := range result {
		// 			todo[ret] = false
		// 		}
		//
		// 		break
		// 	}
		// 	b--
		// }
		//====
		return t
	default:
		return nil
	}
}

// func genIrregularBlock(t *TerminalJson, g graph.Graph, srcIndex int, tgtBlock int, trace *stack.Stack) {
// 	remain := len(t.C) - (t.E-1-tgtBlock)*t.E
// 	tgtRemain := remain - t.E
// 	fmt.Printf("\ntgtBlock: %d, remain: %d, tgtRemain: %d\n", tgtBlock, remain, tgtRemain)
//
// 	// Because of the up-to-down and left-to-right order,
// 	// C[src] is valid to be the first one of target block.
// 	trace.Push(srcIndex)
// 	neighbourOfNeighbours := srcNeighbours(t, g, srcIndex)
// 	for _, neighbour := range neighbourOfNeighbours {
// 		unlink(g, index2id(neighbour), index2id(srcIndex))
// 	}
// 	t.C[srcIndex].B = tgtBlock
// 	remain--
//
// 	rand.Seed(time.Now().Unix())
// 	for remain > tgtRemain {
// 		ok := false
// 		for !ok {
//
// 			currIndex := trace.Peek().(int)
// 			fmt.Printf("\nremain: %v, currIndex: %v, ", remain, currIndex)
//
// 			neighbours := disorderDigits(tgtNeighbours(t, g, currIndex))
// 			fmt.Printf("neighbours: %v, ", neighbours)
//
// 			for _, neighbour := range neighbours {
// 				fmt.Printf("neighbour: %d, ", neighbour)
// 				trace.Push(neighbour)
// 				neighbourOfNeighbours = srcNeighbours(t, g, neighbour)
// 				for _, neighneighbourOfNeighbour := range neighbourOfNeighbours {
// 					unlink(g, index2id(neighneighbourOfNeighbour), index2id(neighbour))
// 				}
// 				t.C[neighbour].B = tgtBlock
//
// 				// as the begin of dfs
// 				var unblockedCellIndex int
// 				for i := len(t.C) - 1; i >= 0; i-- {
// 					if t.C[i].B == 0 {
// 						unblockedCellIndex = i
// 						break
// 					}
// 				}
// 				visited := reachableNum(g, index2id(unblockedCellIndex))
// 				fmt.Printf("From %d reachableNum: %d, ", unblockedCellIndex, visited)
// 				if visited == remain-1 {
// 					ok = true
// 					break
// 				} else {
// 					fmt.Printf("!POP! ")
// 					for _, neighneighbourOfNeighbour := range neighbourOfNeighbours {
// 						link(g, index2id(neighneighbourOfNeighbour), index2id(neighbour))
// 					}
// 					t.C[neighbour].B = 0
// 					trace.Pop()
// 					continue
// 				}
// 			}
// 			if !ok {
// 				fmt.Printf("!!OK! \n")
// 				trace.Pop()
// 			}
// 		}
// 		remain--
// 	}
// }

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
