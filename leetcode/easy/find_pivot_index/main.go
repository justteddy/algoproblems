package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))

	leftSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
	}

	rightSum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1
}
