package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(minProductSum([]int{1, 2}, []int{2, 3}))
	fmt.Println(minProductSum([]int{1, 1, 2}, []int{2, 3, 3}))
}

// naive solution
func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})

	res := 0
	for i, v := range nums1 {
		res += v * nums2[i]
	}

	return res
}
