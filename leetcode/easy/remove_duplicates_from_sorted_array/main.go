package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ptr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[ptr] = nums[i]
			ptr++
		}
	}

	return ptr
}
