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
	return graph.Id(strconv.Itoa(nd.index))
}

func (nd *node) Cell() *Cell {
	return nd.c
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
			nd, err := g.GetNode(graph.Id(index))
			if err != nil {
				nd = newNode(index, &t.C[index])
				g.AddNode(nd)
			}
			// init it's neighbours
			up := index - t.E
			if up > 0 {
				n, err := g.GetNode(graph.Id(up))
				if err != nil {
					n = newNode(up, &t.C[up])
					g.AddNode(n)
				}
				e := graph.NewEdge(1)
				g.AddEdge(nd.Id(), n.Id(), e)
			}
			down := index + t.E
			if down < len(t.C) {
				n, err := g.GetNode(graph.Id(down))
				if err != nil {
					n = newNode(down, &t.C[down])
					g.AddNode(n)
				}
				e := graph.NewEdge(1)
				g.AddEdge(nd.Id(), n.Id(), e)
			}
			left := index - 1
			if left > 0 && t.Row(left) == i {
				n, err := g.GetNode(graph.Id(left))
				if err != nil {
					n = newNode(left, &t.C[left])
					g.AddNode(n)
				}
				e := graph.NewEdge(1)
				g.AddEdge(nd.Id(), n.Id(), e)
			}
			right := index + 1
			if right > 0 && t.Row(right) == i {
				n, err := g.GetNode(graph.Id(right))
				if err != nil {
					n = newNode(right, &t.C[right])
					g.AddNode(n)
				}
				e := graph.NewEdge(1)
				g.AddEdge(nd.Id(), n.Id(), e)
			}
			index++
		}
	}
	return g
}

func Traversal(g graph.Graph, id graph.Id, f func(graph.Node) bool) {
	traversal.Dfs(g, id, f)
}
