package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRecipricolAddAdjacent(t *testing.T) {

	alice := algo.NewVertex("Alice")
	bob := algo.NewVertex("Bob")

	alice.AddAdjacent(bob)

	wantAliceExists := true

	gotAliceExits := false

	for _, vertex := range bob.AdjacentVertices {
		if vertex == alice {
			gotAliceExits = true
		}
	}

	if wantAliceExists != gotAliceExits {
		t.Fatalf("want: %v, got: %v", wantAliceExists, gotAliceExits)
	}

	wantBobExists := true

	gotBobExits := false

	for _, vertex := range alice.AdjacentVertices {
		if vertex == bob {
			gotBobExits = true
		}
	}

	if wantBobExists != gotBobExits {
		t.Fatalf("want: %v, got: %v", wantAliceExists, gotAliceExits)
	}

}

func TestGraphAdd(t *testing.T) {
	t.Parallel()

	want := &algo.Graph{
		Visited: map[*algo.Vertex]interface{}{},
		Vertices: []*algo.Vertex{
			{Value: "Alice"},
			{Value: "Bob"},
		},
	}

	alice := algo.NewVertex("Alice")
	bob := algo.NewVertex("Bob")

	g := algo.NewGraph()
	g.Add(alice)
	g.Add(bob)

	got := g

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestDFS(t *testing.T) {
	t.Parallel()

	alice := algo.NewVertex("Alice")
	bob := algo.NewVertex("Bob")

	alice.AddAdjacent(bob)

}
