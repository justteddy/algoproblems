package main

import "fmt"

// https://leetcode.com/problems/max-consecutive-ones-iii
func main() {
	fmt.Println(longestOnes([]int{1, 1, 0, 0, 1, 1, 1}, 2)) // Output: 5
}

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	zero := 0
	maxLen := 0

	for right < len(nums) {
		if nums[right] == 0 {
			zero++
		}

		for zero > k {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)

		right++
	}

	return maxLen
}
