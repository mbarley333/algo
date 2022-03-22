package algo

import (
	"fmt"
)

const (
	VERTEX_NOT_FOUND = Vertex("vertex not found")
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
	return VERTEX_NOT_FOUND
}

func (g Graph) BFS(startVertex, searchValue Vertex) Vertex {

	if startVertex == searchValue {
		return startVertex
	}

	queue := NewQueue()

	g.visited[startVertex] = struct{}{}
	queue.Enqueue(string(startVertex))

	for len(*queue) > 0 {
		current_vertex, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		}

		if Vertex(current_vertex) == searchValue {
			return Vertex(current_vertex)
		}

		for _, adjacent := range g.adjacent[Vertex(current_vertex)] {
			_, ok := g.visited[adjacent]
			if !ok {
				g.visited[adjacent] = struct{}{}
				queue.Enqueue(string(adjacent))
			}
		}
	}

	return VERTEX_NOT_FOUND
}

type ShortestDistanceGraph struct {
	Verticies                 []*WeightedVertex
	DistanceMap               map[string]int
	PreviousWeightedVertexMap map[string]string
	Unvisited                 []string
	Visited                   map[string]struct{}
}

func NewShortestDistanceGraph() *ShortestDistanceGraph {
	s := &ShortestDistanceGraph{
		DistanceMap:               make(map[string]int),
		PreviousWeightedVertexMap: make(map[string]string),
		Visited:                   make(map[string]struct{}),
	}

	return s
}

func (s *ShortestDistanceGraph) AddVertex(vertex *WeightedVertex) {
	s.Verticies = append(s.Verticies, vertex)
}

func (s ShortestDistanceGraph) GetShortestDistance(startVertex WeightedVertex, endVertex WeightedVertex) []string {

	s.DistanceMap[startVertex.Value] = 0
	current_vertex := startVertex

	for current_vertex.Value != "" {

		s.Visited[current_vertex.Value] = struct{}{}

	}

	return []string{}
}

type WeightedVertex struct {
	Value    string
	Adjacent map[*WeightedVertex]int
}

func NewWeightedVertex(value string) *WeightedVertex {

	w := &WeightedVertex{
		Value:    value,
		Adjacent: make(map[*WeightedVertex]int),
	}
	return w
}

func (w *WeightedVertex) AddAdjacentVertex(vertex *WeightedVertex, weight int) {
	w.Adjacent[vertex] = weight
}
