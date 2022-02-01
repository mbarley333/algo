package algo_test

import (
	"algo"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertAtEnd(t *testing.T) {

	t.Parallel()

	d := algo.DoublyLinkedList{}

	d.InsertAtEnd("first")

	want := algo.DoublyLinkedList{
		FirstNode: &algo.DNode{Data: "first"},
		LastNode:  &algo.DNode{Data: "first"},
	}

	got := d

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	wantLastNode := &algo.DNode{
		Data: "second",
		PreviousNode: &algo.DNode{
			Data: "first",
			NextNode: &algo.DNode{
				Data: "second",
			},
		},
		NextNode: nil,
	}

	d.InsertAtEnd("second")
	// d.AddEndNodeDLL("second")

	gotLastNode := d.LastNode

	fmt.Println(d.LastNode)

	if !cmp.Equal(wantLastNode, gotLastNode) {
		t.Error(cmp.Diff(wantLastNode, gotLastNode))
	}

}
