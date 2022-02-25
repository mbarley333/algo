package algo

import (
	"fmt"
)

type Vertex string

type Graph struct {
	adjacent map[Vertex][]Vertex
	visited  map[Vertex]struct{}
}

func NewGraph() *Graph {
	g := &Graph{
		adjacent: make(map[Vertex][]Vertex),
		visited:  make(map[Vertex]struct{}),
	}
	return g
}

func (g *Graph) AddVertex(vertex Vertex) {
	g.adjacent[vertex] = nil
}

func (g *Graph) AddEdge(vertex1, vertex2 Vertex) error {
	if _, ok := g.adjacent[vertex1]; !ok {
		return fmt.Errorf("unable to create edge: vertex %v does not exist", vertex1)
	}

	g.adjacent[vertex1] = append(g.adjacent[vertex1], vertex2)
	g.adjacent[vertex2] = append(g.adjacent[vertex2], vertex1)
	return nil
}

func (g Graph) DFSrecursive(startVertex, searchValue Vertex) Vertex {

	if startVertex == searchValue {
		return startVertex
	}

	g.visited[startVertex] = struct{}{}

	for _, adjacent := range g.adjacent[startVertex] {

		_, ok := g.visited[adjacent]
		if !ok {
			result := g.DFSrecursive(adjacent, searchValue)
			if result == searchValue {
				return result
			}
		}
	}
	return ""
}
