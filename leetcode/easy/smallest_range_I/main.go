package main

import "fmt"

// https://leetcode.com/problems/smallest-range-i/
func main() {
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 2))
}

func smallestRangeI(nums []int, k int) int {
	min, max := 10_000, 0
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	max -= k
	min += k
	if max <= min {
		return 0
	}

	return max - min
}
