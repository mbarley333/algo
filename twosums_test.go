package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoesTwoSumsReturnNumbersThatSumToTarget(t *testing.T) {

	numbers := []int{0, 1, 7, 2}
	target := 3

	want := 3

	result := algo.TwoSums(numbers, target)

	got := 0

	for _, number := range result {
		got += number
	}

	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}

}

func TestDoesTwoSumsReturnAnEmptySlice(t *testing.T) {

	numbers := []int{0, 1, 7, 2}
	target := 45

	want := []int{}

	got := algo.TwoSums(numbers, target)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
