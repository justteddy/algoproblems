package main

import "fmt"

// https://leetcode.com/problems/permutations/
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	permutations(0, nums, &result)
	return result
}

func permutations(index int, nums []int, result *[][]int) {
	if index == len(nums)-1 {
		*result = append(*result, nums)
		return
	}

	for i := index; i < len(nums); i++ {
		copiedNums := make([]int, len(nums))
		copy(copiedNums, nums)
		copiedNums[index], copiedNums[i] = copiedNums[i], copiedNums[index]
		permutations(index+1, copiedNums, result)
	}
}
