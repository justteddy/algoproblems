package main

import "fmt"

// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/
func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	result := 0
	for _, a1 := range arr1 {
		bad := false
		for _, a2 := range arr2 {
			if abs(a1-a2) <= d {
				bad = true
				break
			}
		}
		if !bad {
			result++
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
