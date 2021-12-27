package main

import "fmt"

// https://leetcode.com/problems/sign-of-the-product-of-an-array/
func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
}

func arraySign(nums []int) int {
	negative := 0
	for _, num := range nums {
		if num < 0 {
			negative++
		}
		if num == 0 {
			return 0
		}
	}

	if negative&1 == 1 {
		return -1
	}

	return 1
}
