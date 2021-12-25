package main

import "fmt"

// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/
func main() {
	fmt.Println(minStartValue([]int{1, 2}))
}

func minStartValue(nums []int) int {
	minPrefixSum := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		prev += nums[i]
		if prev < minPrefixSum {
			minPrefixSum = prev
		}
	}

	if minPrefixSum > 0 {
		return 1
	}

	return 1 - minPrefixSum
}
