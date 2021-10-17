package main

import "fmt"

// https://leetcode.com/problems/house-robber/
func main() {
	fmt.Println(rob([]int{1, 3, 5, 6, 1, 23, 6, 2, 3, 6}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	stolen := make([]int, len(nums))
	stolen[0] = nums[0]
	stolen[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		stolen[i] = max(stolen[i-2]+nums[i], stolen[i-1])
	}

	return stolen[len(stolen)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
