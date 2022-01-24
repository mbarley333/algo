package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPartitioning(t *testing.T) {

	t.Parallel()

	s := algo.Sortable{0, 5, 2, 1, 6, 3}

	want := algo.Sortable{0, 1, 2, 3, 6, 5}

	gotPivotIndex := s.Partition(0, 5)

	got := s

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	wantPivotIndex := 3

	if wantPivotIndex != gotPivotIndex {
		t.Fatalf("want: %d, got:%d", want, gotPivotIndex)
	}
}

func TestQuicksort(t *testing.T) {
	t.Parallel()

	s := algo.Sortable{0, 5, 2, 1, 6, 3}

	s.Quicksort(0, len(s)-1)

	want := algo.Sortable{0, 1, 2, 3, 5, 6}

	got := s

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestQuickselect(t *testing.T) {
	t.Parallel()

	s := algo.Sortable{0, 50, 20, 10, 60, 30}

	want := 10

	got := s.Quickselect(1, 0, len(s)-1)

	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}

}
