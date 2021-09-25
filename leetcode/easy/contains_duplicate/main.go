package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/contains-duplicate/
func main() {
	fmt.Println(containsDuplicate([]int{1, 3, 5, 1, 2, 9}))
}

// map solution
func containsDuplicate(nums []int) bool {
	hm := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := hm[num]; ok {
			return true
		}
		hm[num] = struct{}{}
	}
	return false
}

// sorting soluction
func containsDuplicate_(nums []int) bool {
	sort.Ints(nums)
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			return true
		}
		prev = nums[i]
	}
	return false
}
