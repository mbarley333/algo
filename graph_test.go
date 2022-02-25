package algo_test

import (
	"algo"
	"testing"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	g := algo.NewGraph()

	g.AddVertex("alice")
	g.AddVertex("bob")
	g.AddVertex("joker")
	g.AddVertex("lando")
	g.AddVertex("han")

	err := g.AddEdge("alice", "bob")
	if err != nil {
		t.Fatal(err)
	}

	err = g.AddEdge("alice", "lando")
	if err != nil {
		t.Fatal(err)
	}

	err = g.AddEdge("lando", "han")
	if err != nil {
		t.Fatal(err)
	}

	err = g.AddEdge("bob", "joker")
	if err != nil {
		t.Fatal(err)
	}

	want := algo.Vertex("han")

	got := g.DFSrecursive("alice", "han")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}
