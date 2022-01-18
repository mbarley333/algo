package algo

import (
	"errors"
	"fmt"
)

// recursion is a stack LIFO
// Imagine the function you’re writing has already been implemented by someone else.
// Identify the subproblem of the problem.
// See what happens when you call the function on the subproblem and go from there.

func RecursiveSum(nums []int) int {

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

// return a []string with anagrams
func Anagram(str string) []string {

	// base case
	if len(str) == 1 {
		return []string{str}
	}

	// hold the return valules
	collection := []string{}

	// recurse
	substring_anagrams := Anagram(str[1:])

	for _, anagram := range substring_anagrams {

		// insert the newest character into every position
		for index := 0; index <= len(anagram); index++ {

			new := string(anagram[:index]) + string(str[0]) + string(anagram[index:])
			collection = append(collection, new)
		}

	}

	return collection
}

// given a slice of strings get the total characters

func CharacterCounter(str []string) (int, error) {

	// trap zero length in exported func
	if len(str) == 0 {
		return 0, errors.New("str length == 0")
	}

	// call internal func to perform recurse
	return len(string(str[0])) + characterCounter(str[1:]), nil

}

func characterCounter(str []string) int {

	// base
	if len(str) == 1 {
		return len(str[0])
	}

	// len of current string + whatever came before
	return len(string(str[0])) + characterCounter(str[1:])

}

// given a slice of ints
// return a slice with only even values
func ReturnEven(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	evens := ReturnEven(nums[1:])

	fmt.Println(nums)
	if nums[0]%2 == 0 {
		// append to start of slice
		evens = append(evens, 0)
		copy(evens[1:], evens)
		evens[0] = nums[0]
	}

	return evens
}

// triangle numbers
// 1,2,3 + the sum before it
// 1,3,6,10
func Triangle(n int) int {

	// base if 0 nothing to add
	if n == 0 {
		return 0
	}

	// number + sum before it
	return n + Triangle(n-1)
}

// find the first x in a string
func FirstX(str string) int {

	// base is a the first value
	if string(str[0]) == "x" {
		return 0
	} else {
		// need to track the position with 1
		return 1 + FirstX(string(str[1:]))
	}

}

// given a grid, find all unique paths
func UniquePath(row, col int) int {

	// base is only one row or column which leaves only one path
	if row == 1 || col == 1 {
		return 1
	}

	// paths from a smaller grids
	return UniquePath(row-1, col) + UniquePath(row, col-1)
}

// too many recursions
func MaxTooMany(nums []int) int {

	fmt.Printf("recurse %v\n", nums)
	// base
	if len(nums) == 1 {
		return nums[0]
	}

	// avoid using recursion in conditional statements
	if nums[0] > Max(nums[1:]) {
		return nums[0]
	} else {
		return Max(nums[1:])
	}

}

func Max(nums []int) int {

	fmt.Printf("recurse %v\n", nums)
	// base
	if len(nums) == 1 {
		return nums[0]
	}

	// intermediate variable to solve multiple recursions
	currentMax := Max(nums[1:])

	// recurse
	if nums[0] > currentMax {
		return nums[0]
	} else {
		return currentMax
	}

}
