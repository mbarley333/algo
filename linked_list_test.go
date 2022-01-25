package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLinkedList(t *testing.T) {
	t.Parallel()

	n1 := &algo.Node{
		Data: "once",
	}

	n2 := &algo.Node{
		Data: "upon",
	}

	n3 := &algo.Node{
		Data: "a",
	}

	n4 := &algo.Node{
		Data: "time",
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	l := algo.LinkedList{
		List: n1,
	}

	want := "time"
	got := l.Read(3)

	if want != got {
		t.Fatalf("want:%q, got:%q", want, got)
	}

}

func TestLinkedListSearch(t *testing.T) {
	t.Parallel()

	n1 := &algo.Node{
		Data: "once",
	}

	n2 := &algo.Node{
		Data: "upon",
	}

	n3 := &algo.Node{
		Data: "a",
	}

	n4 := &algo.Node{
		Data: "time",
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	l := algo.LinkedList{
		List: n1,
	}

	want := 0
	got := l.Search("time")

	if want != got {
		t.Fatalf("want:%d, got:%d", want, got)
	}

}

func TestLinkedListInsert(t *testing.T) {
	t.Parallel()

	l := algo.LinkedList{
		List: &algo.Node{
			Data: "yellow",
			Next: &algo.Node{
				Data: "blue",
				Next: &algo.Node{
					Data: "green",
					Next: &algo.Node{
						Data: "red",
						Next: nil}}},
		},
	}

	//want := algo.LinkedList{}
	want := algo.LinkedList{
		List: &algo.Node{
			Data: "yellow",
			Next: &algo.Node{
				Data: "blue",
				Next: &algo.Node{
					Data: "black",
					Next: &algo.Node{
						Data: "green",
						Next: &algo.Node{
							Data: "red",
							Next: nil}}}},
		},
	}

	l.Insert(3, "black")

	got := l

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
