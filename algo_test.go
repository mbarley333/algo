package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBubble(t *testing.T) {
	t.Parallel()

	numbers := []int{4, 2, 7, 1, 3}

	want := []int{1, 2, 3, 4, 7}

	got := algo.Bubble(numbers)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 2, 3, 4, 7}

	type testCase struct {
		target int
		want   bool
		desc   string
	}

	tcs := []testCase{
		{target: 7, want: true, desc: "end"},
		{target: 4, want: true, desc: "four"},
		{target: 1, want: true, desc: "begin"},
		{target: 3, want: true, desc: "mid"},
		{target: 99, want: false, desc: "99"},
	}

	for _, tc := range tcs {

		got := algo.BinarySearch(numbers, tc.target)

		if tc.want != got {
			t.Fatalf("test: %q want: %v, got: %v", tc.desc, tc.want, got)
		}
	}

}

func TestNestedLoopDuplicateSearch(t *testing.T) {
	t.Parallel()

	type testCase struct {
		numbers []int
		want    bool
	}

	tcs := []testCase{
		{numbers: []int{1, 5, 3, 9, 1, 4}, want: true},
		{numbers: []int{1, 5, 3, 9, 17, 4}, want: false},
	}

	for _, tc := range tcs {
		got := algo.NestedLoopDuplicateSearch(tc.numbers)

		if tc.want != got {
			t.Fatalf("want: %v, got: %v", tc.want, got)
		}
	}

}

func TestMapDupilcateSearch(t *testing.T) {

	t.Parallel()

	numbers := []int{1, 5, 3, 9, 1, 4}

	want := true

	got := algo.MapDuplicateSearch(numbers)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

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
