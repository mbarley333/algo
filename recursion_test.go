package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// recursion

func TestRecursiveSum(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}

	want := 15

	got := algo.RecursiveSum(nums)

	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}

func TestRecursiveReverse(t *testing.T) {
	t.Parallel()

	str := "taco"

	want := "ocat"

	// solve aco + t = oca + t at the end
	// RecusiveReverse(str[1:]) + str[0]
	// exit point = i char string

	got := algo.RecursiveReverse(str)

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestCountingX(t *testing.T) {
	t.Parallel()

	str := "abxbxcxd"

	want := 3

	got := algo.CountingX(str)

	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}

}

func TestStaircase(t *testing.T) {

	t.Parallel()

	// 1,1,1
	// 2,1
	// 1,2
	// 3

	n := 3
	want := 4

	got := algo.Staircase(n)

	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}

}

func TestAnagram(t *testing.T) {
	t.Parallel()

	str := "abcd"

	want := []string{
		"abcd",
		"bacd",
		"bcad",
		"bcda",
		"acbd",
		"cabd",
		"cbad",
		"cbda",
		"acdb",
		"cadb",
		"cdab",
		"cdba",
		"abdc",
		"badc",
		"bdac",
		"bdca",
		"adbc",
		"dabc",
		"dbac",
		"dbca",
		"adcb",
		"dacb",
		"dcab",
		"dcba",
	}

	got := algo.Anagram(str)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
