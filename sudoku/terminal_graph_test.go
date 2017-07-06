package sudoku

import (
	"fmt"
	"testing"

	"github.com/jiangyang5157/go-graph/graph"
)

func Test_Traversal(t *testing.T) {
	tJson, _ := Raw2TerminalJson(terminalJson_4x4_3)
	NewGraph(tJson)
	g := NewGraph(tJson)
	nodes := g.Nodes()
	fmt.Printf("Test_Traversal Size of nodes = %d\n", len(nodes))
	visited := 0
	Traversal(g, index2id(0), func(nd graph.Node) bool {
		visited++
		return false
	})
	fmt.Printf("Test_Traversal Visited: %d nodes\n", visited)
}
