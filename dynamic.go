package algo

import "fmt"

// dynamic programming is the process of
// optimizing recursive problems
// that have overlapping subproblems

// Two techniques to solve overlapping
// 1) memoiszation - hash table to store results to make
//    a recursive O(n)

func Fib(num int, memo map[int]int) int {

	if num == 0 || num == 1 {
		return num
	}

	_, ok := memo[num]
	if !ok {
		memo[num] = Fib(num-1, memo) + Fib(num-2, memo)

	}

	return memo[num]
}

// 2) bottom up -- use a loop instead of recursion
func BottomUpFib(n int) int {

	if n == 0 {
		return 0
	}

	a := 0
	b := 1

	for i := 1; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}

	return b
}

func AddUntil100(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	currentSum := AddUntil100(nums[1:])

	// unecessary recursion in conditional logic
	// if nums[0]+AddUntil100(nums[1:]) > 100 {
	// 	return AddUntil100(nums[1:])
	// } else {
	// 	return nums[0] + AddUntil100(nums[1:])
	// }

	// fixed by using intermediate var
	if nums[0]+currentSum > 100 {
		return currentSum
	} else {
		return nums[0] + currentSum
	}

}

func Golomb(num int, memo map[int]int) int {

	if num == 1 {
		return 1
	}

	// memo added
	_, ok := memo[num]
	if !ok {
		memo[num] = 1 + Golomb(num-Golomb((Golomb(num-1, memo)), memo), memo)
	}

	// to many recusive calls
	//return 1 + Golomb(num-Golomb((Golomb(num-1))))

	return memo[num]
}

func UniquePathsMemo(rows, cols int, memo map[string]int) int {

	if rows == 1 || cols == 1 {
		return 1
	}

	str := fmt.Sprintf("%d%d", rows, cols)
	_, ok := memo[str]
	if !ok {
		memo[str] = UniquePathsMemo(rows-1, cols, memo) + UniquePathsMemo(rows, cols-1, memo)
	}

	return memo[str]
}
