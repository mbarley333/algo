package algo_test

import (
	"algo"
	"testing"
)

func TestDoesIsSubsequnceReturnTrue(t *testing.T) {

	parent := []int{1, 2, 3, 5, 9}
	sub := []int{2, 3, 9}

	want := true

	got := algo.IsSubsequence(parent, sub)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestDoesIsSubsequnceReturnFalse(t *testing.T) {

	parent := []int{1, 2, 3, 5, 9}
	sub := []int{2, 3, 99}

	want := false

	got := algo.IsSubsequence(parent, sub)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}
