package sudoku

import "github.com/jiangyang5157/go-graph/graph"

type Node interface {
	graph.Node
	Cell() *Cell
}

type node struct {
	index int
	c     *Cell
}

func (nd *node) String() string {
	return nd.c.String()
}

func (nd *node) Id() graph.Id {
	return graph.Id(nd.index)
}

func (nd *node) Cell() *Cell {
	return nd.c
}

func NewNode(index int, c *Cell) Node {
	return &node{
		index: index,
		c:     c,
	}
}
