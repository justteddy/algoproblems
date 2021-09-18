package main

import "fmt"

// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/
func main() {
	fmt.Println(minOperations([]int{9, 7, 3, 5, 1, 4, 2, 3}))
}

func minOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}

		incrs := abs(nums[i]-nums[i-1]) + 1
		nums[i] += incrs
		result += incrs
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
