package main

import "fmt"

// https://leetcode.com/problems/binary-search/
func main() {
	fmt.Println([]int{-1, 0, 3, 5, 9, 12})
}

func search(nums []int, target int) int {
	mid, lo, hi := 0, 0, len(nums)-1
	for lo <= hi {
		mid = (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
