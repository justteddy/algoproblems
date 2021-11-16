package main

import "fmt"

// https://leetcode.com/problems/missing-number/
func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	dict := make(map[int]struct{})
	for _, num := range nums {
		dict[num] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := dict[i]; !ok {
			return i
		}
	}

	return 0
}
