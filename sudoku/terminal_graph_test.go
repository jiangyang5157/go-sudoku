package sudoku

import (
	"testing"

	"github.com/jiangyang5157/go-graph/graph"
	"github.com/jiangyang5157/go-graph/graph/traversal"
)

func Test_dfs(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)
	visited := 0
	traversal.Dfs(g, Index2Id(2), func(nd graph.Node) bool {
		visited++
		return false
	})
	if len(g.Nodes()) != 81 || visited != 81 {
		t.Error("Test_dfs Visited count should be 81")
	}
}

func Test_targetNeighbours(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)

	tgtNeighbours0 := tgtNeighbours(terminal, g, 0)
	if tgtNeighbours0[0] != 9 || tgtNeighbours0[1] != 1 {
		t.Error("Test_targetNeighbours 0 should be [9 1]")
	}

	tgtNeighbours1 := tgtNeighbours(terminal, g, 1)
	if tgtNeighbours1[0] != 10 || tgtNeighbours1[1] != 0 || tgtNeighbours1[2] != 2 {
		t.Error("Test_targetNeighbours 1 should be [10 0 2]")
	}

	tgtNeighbours11 := tgtNeighbours(terminal, g, 11)
	if tgtNeighbours11[0] != 2 || tgtNeighbours11[1] != 20 || tgtNeighbours11[2] != 10 || tgtNeighbours11[3] != 12 {
		t.Error("Test_targetNeighbours 11 should be [2 20 10 12]")
	}

	tgtNeighbours17 := tgtNeighbours(terminal, g, 17)
	if tgtNeighbours17[0] != 8 || tgtNeighbours17[1] != 26 || tgtNeighbours17[2] != 16 {
		t.Error("Test_targetNeighbours 17 should be [8 26 16]")
	}
}

func Test_srcNeighbours(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)

	srcNeighbours0 := srcNeighbours(terminal, g, 0)
	if srcNeighbours0[0] != 9 || srcNeighbours0[1] != 1 {
		t.Error("Test_srcNeighbours 0 should be [9 1]")
	}

	srcNeighbours1 := srcNeighbours(terminal, g, 1)
	if srcNeighbours1[0] != 10 || srcNeighbours1[1] != 0 || srcNeighbours1[2] != 2 {
		t.Error("Test_srcNeighbours 1 should be [10 0 2]")
	}

	srcNeighbours11 := srcNeighbours(terminal, g, 11)
	if srcNeighbours11[0] != 2 || srcNeighbours11[1] != 20 || srcNeighbours11[2] != 10 || srcNeighbours11[3] != 12 {
		t.Error("Test_srcNeighbours 11 should be [2 20 10 12]")
	}

	srcNeighbours17 := srcNeighbours(terminal, g, 17)
	if srcNeighbours17[0] != 8 || srcNeighbours17[1] != 26 || srcNeighbours17[2] != 16 {
		t.Error("Test_srcNeighbours 17 should be [8 26 16]")
	}
}

func Test_link_unlink(t *testing.T) {
	terminal := NewTerminalJson(9)
	g := NewGraph(terminal)
	tgtNeighbours11 := tgtNeighbours(terminal, g, 11) // [2 20 10 12]
	srcNeighbours11 := srcNeighbours(terminal, g, 11) // [2 20 10 12]
	if tgtNeighbours11[2] != 10 || srcNeighbours11[2] != 10 {
		t.Error("NewGraph(9) related links are wrong")
	}

	unlink(g, Index2Id(11), Index2Id(10))
	tgtNeighbours11 = tgtNeighbours(terminal, g, 11)
	srcNeighbours11 = srcNeighbours(terminal, g, 11)
	if tgtNeighbours11[2] == 10 || srcNeighbours11[2] != 10 {
		t.Error("NewGraph(9) related links are wrong after unlink: 11-x->10")
	}

	link(g, Index2Id(11), Index2Id(10))
	tgtNeighbours11 = tgtNeighbours(terminal, g, 11)
	srcNeighbours11 = srcNeighbours(terminal, g, 11)
	if tgtNeighbours11[2] != 10 || srcNeighbours11[2] != 10 {
		t.Error("NewGraph(9) related links are wrong after link it back: 11--->10")
	}
}
