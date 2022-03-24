package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortedSquaresReturnsSortedSlice(t *testing.T) {
	t.Parallel()

	data := []int{-3, -2, 1}

	want := []int{1, 4, 9}

	got := algo.SortedSquares(data)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestSortedSquaresOptimizedReturnsSortedSlice(t *testing.T) {
	t.Parallel()

	data := []int{-3, -2, 1}

	want := []int{1, 4, 9}

	got := algo.SortedSquaresOptimized(data)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
