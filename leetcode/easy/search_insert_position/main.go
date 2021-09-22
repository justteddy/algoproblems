package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}

func searchInsert(nums []int, target int) int {
	mid, lo, hi := 0, 0, len(nums)-1
	for lo <= hi {
		mid = (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}
