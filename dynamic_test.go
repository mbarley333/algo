package algo_test

import (
	"algo"
	"testing"
)

func TestFib(t *testing.T) {
	t.Parallel()

	want := 55

	memo := map[int]int{}
	got := algo.Fib(10, memo)

	if want != got {
		t.Fatalf("want: %d, got:%d", want, got)
	}

}

func TestBottomUpFib(t *testing.T) {
	t.Parallel()

	want := 55

	got := algo.BottomUpFib(10)

	if want != got {
		t.Fatalf("want: %d, got:%d", want, got)
	}

}

func TestAddUntil100(t *testing.T) {
	t.Parallel()

	want := 99

	test := []int{5, 5, 5, 45, 5, 10, 99}

	got := algo.AddUntil100(test)

	if want != got {
		t.Fatalf("want: %d, got:%d", want, got)
	}

}

func TestGolomb(t *testing.T) {
	t.Parallel()

	want := 3

	memo := map[int]int{}

	got := algo.Golomb(5, memo)

	if want != got {
		t.Fatalf("want: %d, got:%d", want, got)
	}

}

func TestUniquePathsMemo(t *testing.T) {
	t.Parallel()

	want := 28

	rows := 3
	cols := 7
	memo := map[string]int{}

	got := algo.UniquePathsMemo(rows, cols, memo)

	if want != got {
		t.Fatalf("want: %d, got:%d", want, got)
	}

}
