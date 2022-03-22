package algo

// given a target integer, seach a slice for two numbers that sum up to the target value
// interate through the slice...for loop
// given the first number, perform target - first to get the next number needed
// store value:index as a map
// exit point within for loop
// return the two values that sum up to target

func TwoSums(array []int, target int) []int {

	twoMap := map[int]struct{}{}
	secondNumber := 0

	for index, value := range array {

		if len(array) == index {
			break
		}

		secondNumber = target - value

		if _, ok := twoMap[secondNumber]; ok {
			return []int{value, secondNumber}
		}

		twoMap[value] = struct{}{}

	}

	return []int{}
}
