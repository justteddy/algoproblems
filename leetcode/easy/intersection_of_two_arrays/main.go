package main

import "fmt"

/*
	Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

	Example 1:
	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2]

	Example 2:
	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	Output: [9,4]
	Explanation: [4,9] is also accepted.

	Constraints:
	1 <= nums1.length, nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 1000
*/

func main() {
	fmt.Println(intersection([]int{2, 2, 3, 4, 4, 4, 4, 5}, []int{2, 3}))
}

func intersection(nums1 []int, nums2 []int) []int {
	return doIntersection(sliceToMap(nums1), sliceToMap(nums2))
}

func doIntersection(hm1 map[int]struct{}, hm2 map[int]struct{}) []int {
	if len(hm1) < len(hm2) {
		return intersectionByHashMap(hm1, hm2)
	}
	return intersectionByHashMap(hm2, hm1)
}

func intersectionByHashMap(hm1 map[int]struct{}, hm2 map[int]struct{}) []int {
	result := make([]int, 0)
	for n := range hm1 {
		if _, ok := hm2[n]; ok {
			result = append(result, n)
		}
	}
	return result
}

func sliceToMap(nums []int) map[int]struct{} {
	hm := make(map[int]struct{})
	for _, n := range nums {
		hm[n] = struct{}{}
	}

	return hm
}
