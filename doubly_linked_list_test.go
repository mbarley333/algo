package algo_test

import (
	"algo"
	"testing"
)

func TestFindForward(t *testing.T) {
	t.Parallel()

	d := algo.DoublyLinkedList{}

	d.InsertAtEnd("first")
	d.InsertAtEnd("second")
	d.InsertAtEnd("third")

	want := true

	got, err := d.SearchForward("third")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestFindReverse(t *testing.T) {
	t.Parallel()

	d := algo.DoublyLinkedList{}

	d.InsertAtEnd("first")
	d.InsertAtEnd("second")
	d.InsertAtEnd("third")

	want := true

	got, err := d.SearchReverse("first")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}
