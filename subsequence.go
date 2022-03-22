package algo

/*
A subsequence is a slice of numbers that exists within
a larger slice.  Each item of the subsequence slice must exist
in the larger slice and be in the same order.

Order matters with this problem.


Planning this out:
1) Each integer must in the subsequnce must exist in the larger slice
2) Order matters so the order in the subsequnce must exist in the larger slice
3) Iterate through the larger slice forward only since order matters
4) Iterate through the subsequence slice...for loop
5) track the subsequnce iterator
6) loop needs an exit condition
7) increment the interator when match is found
*/

func IsSubsequence(array []int, sequence []int) bool {

	// interator for sequence
	subIndex := 0

	// move through array forward only
	for _, num := range array {

		// exit condition in for loop
		if len(sequence) == subIndex {
			break
		}

		// way to advance the sub iternator value
		if sequence[subIndex] == num {
			subIndex += 1
		}

	}

	return len(sequence) == subIndex
}
