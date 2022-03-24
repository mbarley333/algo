package algo

import "sort"

func SortedSquares(data []int) []int {

	if len(data) < 1 {
		return nil
	}

	result := make([]int, len(data))

	// naive implementation - loop through data slice and put into result slice and sort at end
	// loop with exit condition
	// check for empty slice
	// sorting algo
	// time: O(logN) since you have to sort at end
	// space: O(N) input  slice

	for index, number := range data {
		result[index] = number * number
	}

	sort.Ints(result)

	return result
}

func SortedSquaresOptimized(data []int) []int {

	// move from outside in (start, end) of data slice
	// will have a results slice moving from right to left
	// the big hint is that the data slice is presorted so you should be able to solve in O(n) time
	// for loop with exit condition for when start, end are pointing to same data set

	if len(data) < 1 {
		return nil
	}

	results := make([]int, len(data))

	start := 0
	end := len(data) - 1

	for i := len(data) - 1; i >= 0; i-- {

		//exit
		if start == end {
			results[0] = data[start] * data[start]

		}

		if Abs(data[start]) >= Abs(data[end]) {
			results[i] = data[start] * data[start]
			start += 1
		} else {
			results[i] = data[end] * data[end]
			end -= 1
		}

	}

	return results

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
