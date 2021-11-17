package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	prev, diff := arr[1], arr[1]-arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-prev != diff {
			return false
		}
		prev = arr[i]
	}

	return true
}
