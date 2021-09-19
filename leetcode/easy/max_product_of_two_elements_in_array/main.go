package main

import "fmt"

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	max, prevMax := 0, 0
	for _, n := range nums {
		if n > max {
			prevMax = max
			max = n
		} else if n > prevMax {
			prevMax = n
		}
	}

	return (max - 1) * (prevMax - 1)
}
