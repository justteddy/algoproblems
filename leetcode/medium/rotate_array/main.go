package main

// https://leetcode.com/problems/rotate-array/
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// extra memory solution
func rotate(nums []int, k int) {
	res := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	copy(nums, res)
}
