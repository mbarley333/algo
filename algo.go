package algo

func Bubble(numbers []int) []int {

	length := len(numbers) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < length; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				sorted = false
			}
		}

		length -= 1

	}
	return numbers
}

func BinarySearch(numbers []int, target int) bool {
	min := 0
	max := len(numbers) - 1

	for min <= max {

		mid := (max + min) / 2

		if numbers[mid] == target {
			return true
		} else if target > numbers[mid] {
			min = mid + 1
		} else if target < numbers[mid] {
			max = mid - 1
		}

	}

	return false
}
