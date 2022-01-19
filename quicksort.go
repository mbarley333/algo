package algo

type Sortable []int

func (s Sortable) Paritition(left, right int) int {

	// rightmost element
	pivot_index := right

	// pivot value
	pivot := s[pivot_index]

	// move right index to left
	right -= 1

	for {

		// move left index to right as long
		// as its value is less than the pivot
		for s[left] < pivot {
			left += 1
		}

		// move right index to the left as long
		// as its value is greater than the pivot
		for s[right] > pivot {
			right -= 1
		}

		// when left >= right break
		if left >= right {
			break
			// swap values
		} else {
			s[left], s[right] = s[right], s[left]
			left += 1
		}

	}
	// final step, swap value of left with pivot index
	s[left], s[pivot_index] = s[pivot_index], s[left]

	// index of pivot
	return left
}
