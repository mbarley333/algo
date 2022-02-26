package algo_test

import (
	"algo"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	g := loadGraph()

	want := algo.Vertex("derek")

	got := g.DFSrecursive("alice", "derek")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestBFS(t *testing.T) {
	t.Parallel()

	g := loadGraph()

	want := algo.Vertex("irena")

	got := g.BFS("alice", "irena")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestWeightedGraphTicketPrice(t *testing.T) {
	t.Parallel()

	atlanta := algo.NewWeightedVertex("Atlanta")
	boston := algo.NewWeightedVertex("Boston")
	chicago := algo.NewWeightedVertex("Chicago")
	denver := algo.NewWeightedVertex("Denver")
	el_paso := algo.NewWeightedVertex("El Paso")

	atlanta.AddAdjacentVertex(boston, 100)
	atlanta.AddAdjacentVertex(denver, 160)
	boston.AddAdjacentVertex(chicago, 120)
	boston.AddAdjacentVertex(denver, 180)
	chicago.AddAdjacentVertex(el_paso, 80)
	denver.AddAdjacentVertex(chicago, 40)
	denver.AddAdjacentVertex(el_paso, 140)

	s := algo.NewShortestDistanceGraph()

	s.AddVertex(atlanta)
	s.AddVertex(boston)
	s.AddVertex(chicago)
	s.AddVertex(denver)
	s.AddVertex(el_paso)

	want := []string{"Atlanta", "Denver", "Chicago", "El Paso"}

	got := s.GetShortestDistance(*atlanta, *el_paso)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func loadGraph() *algo.Graph {
	g := algo.NewGraph()

	g.AddVertex("alice")
	g.AddVertex("bob")
	g.AddVertex("candy")
	g.AddVertex("derek")
	g.AddVertex("elaine")

	g.AddVertex("fred")
	g.AddVertex("helen")

	g.AddVertex("gina")
	g.AddVertex("irena")

	err := g.AddEdge("alice", "bob")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("alice", "candy")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("alice", "derek")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("alice", "elaine")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("bob", "fred")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("fred", "helen")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("helen", "candy")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("derek", "gina")
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge("gina", "irena")
	if err != nil {
		fmt.Println(err)
	}

	return g
}
