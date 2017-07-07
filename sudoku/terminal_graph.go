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
	return index2id(nd.index)
}

func (nd *node) Cell() *Cell {
	return nd.c
}

func index2id(index int) graph.Id {
	return graph.Id(strconv.Itoa(index))
}

func id2index(id graph.Id) int {
	ret, _ := strconv.Atoi(string(id))
	return ret
}

func newNode(index int, c *Cell) Node {
	return &node{
		index: index,
		c:     c,
	}
}

func newGraph(t *TerminalJson) graph.Graph {
	g := graph.NewGraph()
	index := 0
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			indexId := index2id(index)
			nd, err := g.GetNode(indexId)
			if err != nil {
				nd = newNode(index, &t.C[index])
				g.AddNode(nd)
			}
			neighbours := t.neighbours(index)
			for _, neighbour := range neighbours {
				neighbourId := index2id(neighbour)
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

func link(g graph.Graph, srcId graph.Id, tgtId graph.Id) {
	g.AddEdge(srcId, tgtId, graph.NewEdge(0))
}

func unlink(g graph.Graph, srcId graph.Id, tgtId graph.Id) {
	g.DeleteEdge(srcId, tgtId)
}

func targetNeighbours(t *TerminalJson, g graph.Graph, index int) []int {
	var ret []int
	id := index2id(index)
	targets, _ := g.GetTargets(id)
	if targets != nil {
		up, down, left, right := t.Up(index), t.Down(index), t.Left(index), t.Right(index)
		if up != -1 && targets[index2id(up)] != nil {
			ret = append(ret, up)
		}
		if down != -1 && targets[index2id(down)] != nil {
			ret = append(ret, down)
		}
		if left != -1 && targets[index2id(left)] != nil {
			ret = append(ret, left)
		}
		if right != -1 && targets[index2id(right)] != nil {
			ret = append(ret, right)
		}
	}
	return ret
}

func srcNeighbours(t *TerminalJson, g graph.Graph, index int) []int {
	var ret []int
	id := index2id(index)
	srcs, _ := g.GetSources(id)
	if srcs != nil {
		up, down, left, right := t.Up(index), t.Down(index), t.Left(index), t.Right(index)
		if up != -1 && srcs[index2id(up)] != nil {
			ret = append(ret, up)
		}
		if down != -1 && srcs[index2id(down)] != nil {
			ret = append(ret, down)
		}
		if left != -1 && srcs[index2id(left)] != nil {
			ret = append(ret, left)
		}
		if right != -1 && srcs[index2id(right)] != nil {
			ret = append(ret, right)
		}
	}
	return ret
}
