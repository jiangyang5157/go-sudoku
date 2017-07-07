package sudoku

import (
	"fmt"
	"testing"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
)

func Test_dfs(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := newGraph(terminal)
	visited := 0
	traversal.Dfs(g, index2id(2), func(nd graph.Node) bool {
		visited++
		return false
	})
	fmt.Printf("Test_dfs Visited: %d/%d nodes\n", visited, len(g.Nodes()))
}

func Test_targetNeighbours(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := newGraph(terminal)
	fmt.Printf("Test_targetNeighbours %d: %v\n", 0, targetNeighbours(terminal, g, 0))
	fmt.Printf("Test_targetNeighbours %d: %v\n", 1, targetNeighbours(terminal, g, 1))
	fmt.Printf("Test_targetNeighbours %d: %v\n", 17, targetNeighbours(terminal, g, 17))
	fmt.Printf("Test_targetNeighbours %d: %v\n", 11, targetNeighbours(terminal, g, 11))
}

func Test_srcNeighbours(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := newGraph(terminal)
	fmt.Printf("Test_srcNeighbours %d: %v\n", 0, srcNeighbours(terminal, g, 0))
	fmt.Printf("Test_srcNeighbours %d: %v\n", 1, srcNeighbours(terminal, g, 1))
	fmt.Printf("Test_srcNeighbours %d: %v\n", 17, srcNeighbours(terminal, g, 17))
	fmt.Printf("Test_srcNeighbours %d: %v\n", 11, srcNeighbours(terminal, g, 11))
}

func Test_link_unlink(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := newGraph(terminal)
	unlink(g, index2id(11), index2id(10))
	fmt.Printf("Test_unlink %d-x->%d targetNeighbours/srcNeighbours of %d: %v / %v\n",
		11, 10, 11, targetNeighbours(terminal, g, 11), srcNeighbours(terminal, g, 11))
	link(g, index2id(11), index2id(10))
	fmt.Printf("Test_link %d-x->%d targetNeighbours/srcNeighbours of %d: %v / %v\n",
		11, 10, 11, targetNeighbours(terminal, g, 11), srcNeighbours(terminal, g, 11))
}
