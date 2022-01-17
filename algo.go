package algo

import "fmt"

func Bubble(numbers []int) []int {

	// O(n^2)

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

	// O(log N)

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

func NestedLoopDuplicateSearch(numbers []int) bool {

	// O(n^2)
	// requires two slices with acting as the comparison

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j && numbers[i] == numbers[j] {
				return true
			}
		}
	}
	return false

}

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

func Linter(str string) error {

	s := Stack{}
	lintMapOpen := map[string]bool{
		"{": false,
		"(": false,
		"[": false,
	}

	lintMapClose := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for _, c := range str {
		letter := string(c)
		_, ok := lintMapOpen[letter]
		if ok {
			s.Push(letter)
			continue
		}

		openBrace, ok := lintMapClose[letter]
		if ok {
			pop, err := s.Pop()
			if err != nil {
				return err
			}
			if openBrace != pop {
				return fmt.Errorf("mismatching braces")
			}
		}
	}

	if len(s) > 0 {
		return fmt.Errorf("missing closing brace")
	}

	return nil
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, error) {

	if len(*s) == 0 {
		return "", fmt.Errorf("no values in stack.  cannot pop until a value is added to stack")
	}

	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return result, nil
}

func Reverser(str string) string {

	s := Stack{}

	for _, c := range str {
		letter := string(c)
		s.Push(letter)
	}

	value := ""

	for i := 0; i < len(str); i++ {
		pop, err := s.Pop()

		if err != nil {
			fmt.Println(err)
		}

		value += pop
	}
	return value
}

// figure out the return value for the prior subproblem
// assume it will all work out initially
// base case is the exit point

func RecursiveSum(nums []int) int {

	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + RecursiveSum(nums[1:])
}

func RecursiveReverse(str string) string {

	if len(str) == 1 {
		return string(str[0])
	}

	return RecursiveReverse(str[1:]) + string(str[0])
}

func CountingX(str string) int {

	if len(str) == 0 {
		return 0
	}

	if string(str[0]) == "x" {
		return 1 + CountingX(str[1:])
	} else {
		return CountingX(str[1:])
	}

}

func Staircase(n int) int {

	if n == 1 || n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	return Staircase(n-1) + Staircase(n-2) + Staircase(n-3)

}

func Anagram(str string) []string {

	if len(str) == 1 {
		return []string{str}
	}

	collection := []string{}

	substring_anagrams := Anagram(str[1:])

	fmt.Println(str)

	for _, anagram := range substring_anagrams {

		for index := 0; index <= len(anagram); index++ {

			new := string(anagram[:index]) + string(str[0]) + string(anagram[index:])
			collection = append(collection, new)
		}

	}

	return collection
}
