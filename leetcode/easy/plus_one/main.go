package main

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i]++
		return digits
	}
	return append([]int{1}, digits...)
}
