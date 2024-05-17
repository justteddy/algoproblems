package main

import "fmt"

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}))
}

func anagramMappings(nums1 []int, nums2 []int) []int {
	idx := make(map[int]int, len(nums2))
	for k, v := range nums2 {
		idx[v] = k
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = idx[nums1[i]]
	}

	return nums1
}
