package main

import "fmt"

// https://leetcode.com/problems/maximum-subarray/
func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	runningSum, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		runningSum = max(runningSum+v, v)
		if runningSum > sum {
			sum = runningSum
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
