package main

// https://leetcode.com/problems/find-the-difference-of-two-arrays
func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)

	m1 := make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}

	m2 := make(map[int]bool)
	for _, v := range nums2 {
		m2[v] = true
	}

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			res[0] = append(res[0], k)
		}
	}

	for k := range m2 {
		if _, ok := m1[k]; !ok {
			res[1] = append(res[1], k)
		}
	}

	return res
}
