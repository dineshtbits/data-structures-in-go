package graph

import (
	"fmt"

	"github.com/dineshtbits/data-structures-in-go/queue"
	"github.com/dineshtbits/data-structures-in-go/stack"
)

type GraphMatrix struct {
	edges [][]int
}

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

func (g *Graph) visitDfs(id int, visited map[int]bool) {
	visited[id] = true
	fmt.Printf("%v ", id)
	for k := range g.nodes[id].edges {
		if !visited[k] {
			g.visitDfs(k, visited)
		}
	}
}

func (g *Graph) DFS() {
	visited := make(map[int]bool)
	for i := 0; i < len(g.nodes); i++ {
		if !visited[g.nodes[i].id] {
			g.visitDfs(g.nodes[i].id, visited)
		}
	}
}

func (g *Graph) visitBFS(node int, visited map[int]bool) {
	q := queue.Queue{}

	visited[node] = true
	q.Enqueue(node)

	for q.Size() > 0 {
		n := q.Dequeue()
		fmt.Printf("%v", n)

		for k := range g.nodes[n.(int)].edges {
			if !visited[k] {
				visited[k] = true
				q.Enqueue(k)
			}
		}
	}
}

func (g *Graph) BFS() {
	visited := make(map[int]bool)
	for i := 0; i < len(g.nodes); i++ {
		if !visited[g.nodes[i].id] {
			g.visitBFS(g.nodes[i].id, visited)
		}
	}
}

func NewMatrix(nodes int) *GraphMatrix {
	m := &GraphMatrix{edges: [][]int{}}
	for i := 0; i < nodes; i++ {
		row := []int{}
		for j := 0; j < nodes; j++ {
			row = append(row, 0)
		}
		m.edges = append(m.edges, row)
	}
	return m
}

func (g *Graph) AddNode() {
	g.nodes = append(g.nodes, &Node{id: len(g.nodes), edges: make(map[int]int)})
}

func (m *GraphMatrix) AddMatrixNode() {
	row := []int{}
	for i := 0; i < len(m.edges); i++ {
		m.edges[i] = append(m.edges[i], 0)
		row = append(row, 0)
	}
	row = append(row, 0)
	m.edges = append(m.edges, row)
}

func (m *GraphMatrix) AddMatrixEdge(from, to, weight int) {
	m.edges[from][to] = weight
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

func (g *Graph) TopologicalSort() {
	s := &stack.Stack{}
	visited := make(map[int]bool)

	for i := 0; i < len(g.nodes); i++ {
		if !visited[i] {
			g.Tsort(i, s, visited)
		}
	}
	for s.Size() > 0 {
		fmt.Printf("%v ", s.Pop().(int))
	}
}

func (g *Graph) Tsort(n int, s *stack.Stack, visited map[int]bool) {

	visited[n] = true

	for i := 0; i < len(g.nodes[n].edges); i++ {
		if !visited[i] {
			g.Tsort(i, s, visited)
		}
	}
	s.Push(n)
}
