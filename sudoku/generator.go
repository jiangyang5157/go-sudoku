package sudoku

import (
	"math"
	"math/rand"
	"time"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
)

type BlockMode int

const (
	SQUARE BlockMode = iota
	IRREGULAR
)

func GenString(blockMode int, edge int, minSubGiven int, minTotalGiven int) string {
	return string(GenByte(blockMode, edge, minSubGiven, minTotalGiven))
}

func GenByte(blockMode int, edge int, minSubGiven int, minTotalGiven int) []byte {
	t := GenTerminalJson(blockMode, edge, minSubGiven, minTotalGiven)
	ret, _ := TerminalJson2Raw(t)
	return ret
}

func GenTerminalJson(blockMode int, edge int, minSubGiven int, minTotalGiven int) *TerminalJson {
	rand.Seed(time.Now().Unix())

	t := NewTerminalJson(edge)
	t = t.genBlock(blockMode)
	var ret *TerminalJson
	ok := false
	for !ok {
		ret = t.Clone()
		ret = ret.genMaterial(blockMode)
		// valid material must have at least one solution
		ret = SolveTerminalJson(ret)
		if ret != nil {
			ok = true
		}
	}

	ret = ret.genPuzzle(minSubGiven, minTotalGiven)
	return ret
}

func (t *TerminalJson) genBlock(blockMode int) *TerminalJson {
	// Gen diagonal square
	square := int(math.Sqrt(float64(t.E)))
	for i := 0; i < len(t.C); i++ {
		row, col := t.Row(i), t.Col(i)
		c := &t.C[i]
		c.B = (row/square)*square + col/square
	}

	mode := BlockMode(blockMode)
	switch mode {
	case SQUARE:
		return t
	case IRREGULAR:
		if t.E == 1 {
			// 1*1 sudoku doesn't require swap
			return t
		}

		g := NewGraph(t)
		// unlink between blocks
		for i := 0; i < t.E; i++ {
			for j := 0; j < t.E; j++ {
				if i < t.E-1 && (i+1)%square == 0 {
					top, bottom := Index2Id(t.Index(i, j)), Index2Id(t.Index(i+1, j))
					unlink(g, top, bottom)
					unlink(g, bottom, top)
				}
				if j < t.E-1 && (j+1)%square == 0 {
					left, right := Index2Id(t.Index(i, j)), Index2Id(t.Index(i, j+1))
					unlink(g, left, right)
					unlink(g, right, left)
				}
			}
		}

		attemptsCount, attemptsMin, attemptsMax := 0, len(t.C)/square, 1000
		for i := 0; i < attemptsMin; i++ {
			attemptsCount++
			if !swap(t, g) {
				i--
			}
			if attemptsCount >= attemptsMax {
				break
			}
		}
		return t
	default:
		return nil
	}
}

func swap(t *TerminalJson, g graph.Graph) bool {
	// Gen random aIndex and bIndex
	aIndex, bIndex := -1, -1
	aIndex = rand.Intn(len(t.C))
	aNbs := t.Neighbours(aIndex)
	aNbs = disorderDigits(aNbs)
	for _, aNb := range aNbs {
		if t.C[aNb].B != t.C[aIndex].B {
			random := rand.Intn(t.E)
			traversal.Dfs(g, Index2Id(aNb), func(nd graph.Node) bool {
				bIndex = Id2Index(nd.Id())
				random--
				return random < 0
			})
			break
		}
	}
	if aIndex == -1 || bIndex == -1 {
		return false
	}

	// Swap aIndex and bIndex
	aBlock, bBlock := t.C[aIndex].B, t.C[bIndex].B
	aIndexId, bIndexId := Index2Id(aIndex), Index2Id(bIndex)
	bNbs := t.Neighbours(bIndex)

	t.C[aIndex].B, t.C[bIndex].B = bBlock, aBlock
	for _, aNb := range aNbs {
		aNbId := Index2Id(aNb)
		if t.C[aNb].B == aBlock {
			unlink(g, aIndexId, aNbId)
			unlink(g, aNbId, aIndexId)
		}
		if t.C[aNb].B == bBlock {
			link(g, aIndexId, aNbId)
			link(g, aNbId, aIndexId)
		}
	}
	for _, bNb := range bNbs {
		bNbId := Index2Id(bNb)
		if t.C[bNb].B == bBlock {
			unlink(g, bIndexId, bNbId)
			unlink(g, bNbId, bIndexId)
		}
		if t.C[bNb].B == aBlock {
			link(g, bIndexId, bNbId)
			link(g, bNbId, bIndexId)
		}
	}

	// Validate
	aValidation, bValidation := 0, 0
	traversal.Dfs(g, aIndexId, func(nd graph.Node) bool {
		bValidation++
		return false
	})
	traversal.Dfs(g, bIndexId, func(nd graph.Node) bool {
		aValidation++
		return false
	})

	if aValidation != t.E || bValidation != t.E {
		// Undo swap
		t.C[aIndex].B, t.C[bIndex].B = aBlock, bBlock
		for _, aNb := range aNbs {
			aNbId := Index2Id(aNb)
			if t.C[aNb].B == aBlock {
				link(g, aIndexId, aNbId)
				link(g, aNbId, aIndexId)
			}
			if t.C[aNb].B == bBlock {
				unlink(g, aIndexId, aNbId)
				unlink(g, aNbId, aIndexId)
			}
		}
		for _, bNb := range bNbs {
			bNbId := Index2Id(bNb)
			if t.C[bNb].B == bBlock {
				link(g, bIndexId, bNbId)
				link(g, bNbId, bIndexId)
			}
			if t.C[bNb].B == aBlock {
				unlink(g, bIndexId, bNbId)
				unlink(g, bNbId, bIndexId)
			}
		}
		return false
	}

	return true
}

func (t *TerminalJson) genMaterial(blockMode int) *TerminalJson {
	mode := BlockMode(blockMode)
	switch mode {
	case SQUARE:
		// Fill diagonal square by random digits
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
	case IRREGULAR:
		// TODO: filling
		return t
	default:
		return nil
	}
}

func (t *TerminalJson) genPuzzle(minSubGiven int, minTotalGiven int) *TerminalJson {
	s := newSudoku(t)

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

// ...
// //rand.Seed(time.Now().Unix())
// var ret *TerminalJson
// ok := false
// for !ok {
// 	ret = t.Clone()
// 	b := ret.E - 1
// 	g := NewGraph(ret)
// 	// up-to-down and left-to-right
// 	for i := 0; i < len(ret.C); i++ {
// 		if b == 0 {
// 			// rest of cells belongs to block 0
// 			ok = true
// 			break
// 		}
// 		if ret.C[i].B > 0 {
// 			// already be assigned to a particular block
// 			continue
// 		}
// 		if genIrregularBlock(ret, g, b, i) {
// 			b--
// 		} else {
// 			// retry
// 			ok = false
// 			break
// 		}
// 	}
// }
// ...
// func genIrregularBlock(t *TerminalJson, g graph.Graph, block int, begin int) bool {
// 	remain := len(t.C) - (t.E-1-block)*t.E
// 	tgtRemain := remain - t.E
//
// 	// C[begin] is valid to be the first one of target block.
// 	trace := stack.NewStack()
// 	nbsOfBegin := srcNeighbours(t, g, begin)
// 	for _, nb := range nbsOfBegin {
// 		unlink(g, Index2Id(nb), Index2Id(begin))
// 	}
// 	trace.Push(begin)
// 	t.C[begin].B = block
// 	remain--
//
// 	for remain > tgtRemain {
// 		ok := false
// 		for !ok {
// 			if trace.IsEmpty() {
// 				break
// 			}
// 			index := trace.Peek().(int)
// 			nbs := disorderDigits(tgtNeighbours(t, g, index))
//
// 			for _, nb := range nbs {
// 				trace.Push(nb)
// 				t.C[nb].B = block
// 				nbsOfNbs := srcNeighbours(t, g, nb)
// 				for _, nbOfNbs := range nbsOfNbs {
// 					unlink(g, Index2Id(nbOfNbs), Index2Id(nb))
// 				}
//
// 				// dow-to-up and right-to-left
// 				var indexOfAnyBlock0 int
// 				for i := len(t.C) - 1; i >= 0; i-- {
// 					if t.C[i].B == 0 {
// 						indexOfAnyBlock0 = i
// 						break
// 					}
// 				}
//
// 				visited := 0
// 				traversal.Dfs(g, Index2Id(indexOfAnyBlock0), func(nd graph.Node) bool {
// 					visited++
// 					return false
// 				})
//
// 				if visited == remain-1 {
// 					ok = true
// 					break
// 				} else {
// 					for _, nbOfNbs := range nbsOfNbs {
// 						link(g, Index2Id(nbOfNbs), Index2Id(nb))
// 					}
// 					t.C[nb].B = 0
// 					trace.Pop()
// 				}
// 			}
// 			if !ok {
// 				trace.Pop()
// 			}
// 		}
// 		remain--
// 	}
// 	return !trace.IsEmpty()
// }
// ....
// func reachableNum(g graph.Graph, id graph.Id) int {
// 	visited := 0
// 	traversal.Dfs(g, id, func(nd graph.Node) bool {
// 		visited++
// 		return false
// 	})
// 	return visited
// }
