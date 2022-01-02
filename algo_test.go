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
