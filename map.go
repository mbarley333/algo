package algo

func MapDuplicateSearch(numbers []int) bool {

	// O(n)
	// alternative to double loop

	duplicateMap := map[int]bool{}

	for _, number := range numbers {

		_, ok := duplicateMap[number]

		if ok {
			return true
		} else {
			duplicateMap[number] = true
		}
	}

	return false

}
