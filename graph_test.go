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
	g.AddVertex("lando")

	err := g.AddEdge("alice", "bob")
	if err != nil {
		t.Fatal(err)
	}

	err = g.AddEdge("alice", "lando")
	if err != nil {
		t.Fatal(err)
	}

	want := algo.Vertex("bob")

	got := g.DFSrecursive("alice", "bob")

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}
