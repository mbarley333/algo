package algo_test

import (
	"algo"
	"testing"
)

func TestMapDupilcateSearch(t *testing.T) {

	t.Parallel()

	numbers := []int{1, 5, 3, 9, 1, 4}

	want := true

	got := algo.MapDuplicateSearch(numbers)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}
