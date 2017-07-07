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

func NewGraph(t *TerminalJson) graph.Graph {
	g := graph.NewGraph()
	index := 0
	for i := 0; i < t.E; i++ {
		for j := 0; j < t.E; j++ {
			// init the node
			indexId := index2id(index)
			nd, err := g.GetNode(indexId)
			if err != nil {
				nd = newNode(index, &t.C[index])
				g.AddNode(nd)
			}
			// init it's neighbours
			neighbours := genNeighbours(t, index)
			for _, neighbour := range neighbours {
				addEdge(t, g, index, neighbour)
			}
			index++
		}
	}
	return g
}

func genNeighbours(t *TerminalJson, srcIndex int) []int {
	var ret []int
	up, down, left, right := t.Up(srcIndex), t.Down(srcIndex), t.Left(srcIndex), t.Right(srcIndex)
	if up != -1 {
		ret = append(ret, up)
	}
	if down != -1 {
		ret = append(ret, down)
	}
	if left != -1 {
		ret = append(ret, left)
	}
	if right != -1 {
		ret = append(ret, right)
	}
	return ret
}

func addEdge(t *TerminalJson, g graph.Graph, srcIndex int, tgtIndex int) {
	tgtId := index2id(tgtIndex)
	n, err := g.GetNode(tgtId)
	if err != nil {
		n = newNode(tgtIndex, &t.C[tgtIndex])
		g.AddNode(n)
	}
	e := graph.NewEdge(0)
	g.AddEdge(index2id(srcIndex), tgtId, e)
}

func Traversal(g graph.Graph, id graph.Id, f func(graph.Node) bool) {
	traversal.Dfs(g, id, f)
}
