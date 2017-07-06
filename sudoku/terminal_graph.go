package sudoku

import (
	"fmt"
	"strconv"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
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

func newNode(index int, c *Cell) Node {
	return &node{
		index: index,
		c:     c,
	}
}

func NewGraph(t *TerminalJson) graph.Graph {
	g := graph.NewGraph()
	index := 0
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			// init the node
			nd, err := g.GetNode(index2id(index))
			if err != nil {
				nd = newNode(index, &t.C[index])
				g.AddNode(nd)
			}
			// init it's neighbours
			up, down, left, right := index-t.E, index+t.E, index-1, index+1
			if up > 0 {
				addEdge(t, g, nd, up)
			}
			if down < len(t.C) {
				addEdge(t, g, nd, down)
			}
			if left > 0 && t.Row(left) == i {
				addEdge(t, g, nd, left)
			}
			if right < len(t.C) && t.Row(right) == i {
				addEdge(t, g, nd, right)
			}
			index++
		}
	}
	return g
}

func addEdge(t *TerminalJson, g graph.Graph, src graph.Node, tgtIndex int) {
	n, err := g.GetNode(index2id(tgtIndex))
	if err != nil {
		n = newNode(tgtIndex, &t.C[tgtIndex])
		g.AddNode(n)
	}
	e := graph.NewEdge(0)
	g.AddEdge(src.Id(), n.Id(), e)
}

func Traversal(g graph.Graph, id graph.Id, f func(graph.Node) bool) {
	traversal.Dfs(g, id, f)
}
