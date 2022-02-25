package algo_test

import (
	"algo"
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	g := LoadGraph()

	want := algo.Vertex("derek")

	got := g.DFSrecursive("alice", "derek")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestBFS(t *testing.T) {
	t.Parallel()

	g := LoadGraph()

	want := algo.Vertex("irena")

	got := g.BFS("alice", "irena")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func LoadGraph() *algo.Graph {
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
