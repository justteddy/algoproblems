package majority_element

import "sort"

/*
	Given an array nums of size n, return the majority element.
	The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

	Example 1:
	Input: nums = [3,2,3]
	Output: 3

	Example 2:
	Input: nums = [2,2,1,1,1,2,2]
	Output: 2
*/

// sorting solution
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// hashmap solution
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
			continue
		}
		m[num]++
	}

	var max, maxNum int
	for num, count := range m {
		if count > max {
			max = count
			maxNum = num
		}
	}

	return maxNum
}
