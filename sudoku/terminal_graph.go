package sudoku

import (
	"fmt"
	"strconv"

	"github.com/jiangyang5157/go-graph/graph"
)

type Node interface {
	graph.Node
	Cell() *Cell
}

type node struct {
	index int
	c     *Cell
}

func (nd *node) String() string {
	return fmt.Sprintf("%s(%s)", nd.c.String(), nd.Id())
}

func (nd *node) Id() graph.Id {
	return Index2Id(nd.index)
}

func (nd *node) Cell() *Cell {
	return nd.c
}

func Index2Id(index int) graph.Id {
	return graph.Id(strconv.Itoa(index))
}

func Id2Index(id graph.Id) int {
	ret, _ := strconv.Atoi(string(id))
	return ret
}

func NewGraph(t *TerminalJson) graph.Graph {
	g := graph.NewGraph()
	index := 0
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			indexId := Index2Id(index)
			nd, err := g.GetNode(indexId)
			if err != nil {
				nd = newNode(index, &t.C[index])
				g.AddNode(nd)
			}
			neighbours := t.Neighbours(index)
			for _, neighbour := range neighbours {
				neighbourId := Index2Id(neighbour)
				n, err := g.GetNode(neighbourId)
				if err != nil {
					n = newNode(neighbour, &t.C[neighbour])
					g.AddNode(n)
				}
				link(g, indexId, neighbourId)
			}
			index++
		}
	}
	return g
}

func newNode(index int, c *Cell) Node {
	return &node{
		index: index,
		c:     c,
	}
}

func link(g graph.Graph, srcId graph.Id, tgtId graph.Id) {
	g.AddEdge(srcId, tgtId, graph.NewEdge(0))
}

func unlink(g graph.Graph, srcId graph.Id, tgtId graph.Id) {
	g.DeleteEdge(srcId, tgtId)
}

func tgtNeighbours(t *TerminalJson, g graph.Graph, index int) []int {
	var ret []int
	id := Index2Id(index)
	targets, _ := g.GetTargets(id)
	if targets != nil {
		up, down, left, right := t.Up(index), t.Down(index), t.Left(index), t.Right(index)
		if up != -1 && targets[Index2Id(up)] != nil {
			ret = append(ret, up)
		}
		if down != -1 && targets[Index2Id(down)] != nil {
			ret = append(ret, down)
		}
		if left != -1 && targets[Index2Id(left)] != nil {
			ret = append(ret, left)
		}
		if right != -1 && targets[Index2Id(right)] != nil {
			ret = append(ret, right)
		}
	}
	return ret
}

func srcNeighbours(t *TerminalJson, g graph.Graph, index int) []int {
	var ret []int
	id := Index2Id(index)
	srcs, _ := g.GetSources(id)
	if srcs != nil {
		up, down, left, right := t.Up(index), t.Down(index), t.Left(index), t.Right(index)
		if up != -1 && srcs[Index2Id(up)] != nil {
			ret = append(ret, up)
		}
		if down != -1 && srcs[Index2Id(down)] != nil {
			ret = append(ret, down)
		}
		if left != -1 && srcs[Index2Id(left)] != nil {
			ret = append(ret, left)
		}
		if right != -1 && srcs[Index2Id(right)] != nil {
			ret = append(ret, right)
		}
	}
	return ret
}
