package algo_test

import (
	"algo"
	"testing"
)

// stack

func TestLinter(t *testing.T) {

	t.Parallel()

	str := "(var x = {y: [1,2,3]})"

	err := algo.Linter(str)
	if err != nil {
		t.Fatal(err)
	}

}

func TestReverserStack(t *testing.T) {

	t.Parallel()

	str := "taco"

	want := "ocat"

	got := algo.Reverser(str)

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}

}
