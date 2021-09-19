package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hm := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		hm[v] = i
	}

	for i, v := range nums1 {
		var found bool
		for j := hm[v]; j < len(nums2); j++ {
			if nums2[j] > v {
				nums1[i] = nums2[j]
				found = true
				delete(hm, v)
				break
			}
		}

		if !found {
			nums1[i] = -1
		}
	}

	return nums1
}
