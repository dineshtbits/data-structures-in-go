package graph

import "fmt"

type Graph struct {
	nodes []*Node
}

type Node struct {
	id    int
	edges map[int]int
}

func New() *Graph {
	return &Graph{nodes: []*Node{}}
}

func (g *Graph) AddNode() {
	g.nodes = append(g.nodes, &Node{id: len(g.nodes), edges: make(map[int]int)})
}

func (g *Graph) ShowNodes() {
	for i := 0; i < len(g.nodes); i++ {
		fmt.Printf("%v", g.nodes[i].id)
	}
}

func (g *Graph) AddEdge(from, to, w int) {
	g.nodes[from].edges[to] = w
}

func (g *Graph) ShowEdges() [][]int {
	edges := [][]int{}
	for i := 0; i < len(g.nodes); i++ {
		for n, w := range g.nodes[i].edges {
			edges = append(edges, []int{i, n, w})
		}
	}
	return edges
}
