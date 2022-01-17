package algo

import "fmt"

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
