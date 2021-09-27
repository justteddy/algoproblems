package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/
func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersection(nums1, nums2)
	}
	return intersection(nums2, nums1)
}

func intersection(primary, secondary []int) []int {
	m := make(map[int]int)
	for _, v := range primary {
		m[v]++
	}

	result := make([]int, 0)
	for _, v := range secondary {
		if count, ok := m[v]; ok && count != 0 {
			result = append(result, v)
			m[v]--
		}
	}

	return result
}
