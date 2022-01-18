package algo_test

import (
	"algo"
	"testing"
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

	want := "a"
	got := l.Read(2)

	if want != got {
		t.Fatalf("want:%q, got:%q", want, got)
	}

}
