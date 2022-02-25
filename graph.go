package algo

import (
	"fmt"
)

type Vertex string

type Graph struct {
	Adjacent map[Vertex][]Vertex
	Visited  map[Vertex]struct{}
}

func NewGraph() *Graph {
	g := &Graph{
		Adjacent: make(map[Vertex][]Vertex),
		Visited:  make(map[Vertex]struct{}),
	}
	return g
}

func (g *Graph) AddVertex(vertex Vertex) {
	g.Adjacent[vertex] = nil
}

func (g *Graph) AddEdge(vertex1, vertex2 Vertex) error {
	if _, ok := g.Adjacent[vertex1]; !ok {
		return fmt.Errorf("unable to create edge: vertex %v does not exist", vertex1)
	}

	g.Adjacent[vertex1] = append(g.Adjacent[vertex1], vertex2)
	g.Adjacent[vertex2] = append(g.Adjacent[vertex2], vertex1)
	return nil
}

func (g Graph) DFSrecursive(startVertex, searchValue Vertex) Vertex {

	if startVertex == searchValue {
		return startVertex
	}

	g.Visited[startVertex] = struct{}{}

	for _, adjacent := range g.Adjacent[startVertex] {

		_, ok := g.Visited[adjacent]
		if !ok {
			result := g.DFSrecursive(adjacent, searchValue)
			if result == searchValue {
				return result
			}

		}
	}
	return ""
}
