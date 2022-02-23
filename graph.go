package algo

type Vertex struct {
	Value            string
	AdjacentVertices []*Vertex
}

func NewVertex(value string) *Vertex {
	v := &Vertex{
		Value: value,
	}
	return v
}

// reciprocol add to adjacent
func (v *Vertex) AddAdjacent(vertex *Vertex) {

	// bail if entry exists
	for _, exisitingVertex := range v.AdjacentVertices {
		if v.Value == exisitingVertex.Value {
			return
		}
	}

	v.AdjacentVertices = append(v.AdjacentVertices, vertex)

	// add adjacent to other side too
	vertex.AdjacentVertices = append(vertex.AdjacentVertices, v)

}

type Graph struct {
	Vertices []*Vertex
	Visited  map[*Vertex]interface{}
}

func NewGraph() *Graph {
	g := &Graph{
		Visited: make(map[*Vertex]interface{}),
	}
	return g
}

func (g *Graph) Add(vertex *Vertex) {
	g.Vertices = append(g.Vertices, vertex)
}
