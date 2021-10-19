package main

import "fmt"

// https://leetcode.com/problems/sum-of-unique-elements/
func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 4, 4, 5, 6, 7, 7, 7, 7, 8}))
}

func sumOfUnique(nums []int) int {
	hm := make(map[int]int)
	result := 0
	for _, num := range nums {
		hm[num]++
		switch hm[num] {
		case 2:
			result -= num
		case 1:
			result += num
		}
	}
	return result
}
