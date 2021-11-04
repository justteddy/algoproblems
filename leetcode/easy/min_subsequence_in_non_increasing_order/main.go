package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/
func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
}

func minSubsequence(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	subSum := 0
	for i := 0; i < len(nums); i++ {
		subSum += nums[i]
		if subSum > sum/2 {
			return nums[:i+1]
		}
	}
	return nil
}

// solution, for case, when you can check only subsequences, not every comnbination of array elements
// stupid task, stupid condition
func minSubsequence_(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	maxSubSum, minIndices := 0, make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		subSum := nums[i]
		indices := []int{i}
		for j := i + 1; j < len(nums); j++ {
			if subSum > sum/2 {
				break
			}
			subSum += nums[j]
			indices = append(indices, j)
		}
		if subSum <= sum/2 {
			continue
		}
		if len(indices) < len(minIndices) || (len(indices) == len(minIndices) && subSum > maxSubSum) {
			maxSubSum = subSum
			minIndices = indices
		}
	}
	result := nums[minIndices[0] : minIndices[len(minIndices)-1]+1]
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
