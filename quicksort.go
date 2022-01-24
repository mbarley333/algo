package algo

type Sortable []int

func (s Sortable) Partition(left, right int) int {

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

func (s *Sortable) Quicksort(left, right int) {

	if right-left <= 0 {
		return
	}

	pivot_index := s.Partition(left, right)

	s.Quicksort(left, pivot_index-1)

	s.Quicksort(pivot_index+1, right)
}

func (s *Sortable) Quickselect(kth_lowest, left, right int) int {

	if right-left == 0 {
		return (*s)[left]
	}

	pivot_index := s.Partition(left, right)

	if kth_lowest < pivot_index {
		return s.Quickselect(kth_lowest, left, pivot_index-1)
	} else if kth_lowest > pivot_index {
		return s.Quickselect(kth_lowest, pivot_index+1, right)
	} else {
		return (*s)[pivot_index]
	}

}
