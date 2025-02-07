package main

import "fmt"

// https://leetcode.com/problems/maximum-ascending-subarray-sum/
func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}

func maxAscendingSum(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currSum += nums[i]
		} else {
			currSum = nums[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
