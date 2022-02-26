package algo_test

import (
	"algo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	q := algo.NewQueue()

	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	want := []string{"a", "b", "c"}
	result := []string{}

	for len(*q) > 0 {
		r, err := q.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		result = append(result, r)

	}

	got := result

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
