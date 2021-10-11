package main

import "fmt"

// https://leetcode.com/problems/sort-array-by-parity-ii/
func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

// time optimized solution
func sortArrayByParityII(nums []int) []int {
	result := make([]int, len(nums))

	idxEven, idxOdd := 0, 1
	for _, num := range nums {
		if num&1 == 0 {
			result[idxEven] = num
			idxEven += 2
		} else {
			result[idxOdd] = num
			idxOdd += 2
		}
	}

	return result
}

// memory optimized solution
func sortArrayByParityII_(nums []int) []int {
	for slow := 0; slow < len(nums)-1; slow++ {
		if slow&1 == nums[slow]&1 {
			continue
		}
		for fast := slow + 1; fast < len(nums); fast++ {
			if nums[fast]&1 == slow&1 {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				break
			}
		}
	}

	return nums
}
