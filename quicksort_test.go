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

	gotPivotIndex := s.Paritition(0, 5)

	got := s

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	wantPivotIndex := 3

	if wantPivotIndex != gotPivotIndex {
		t.Fatalf("want: %d, got:%d", want, gotPivotIndex)
	}
}
