package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/mean-of-array-after-removing-some-elements/
func main() {
	fmt.Println(trimMean([]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}))
}

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	cutLen := float64(len(arr)) / 100 * 5
	sum := 0

	for i := int(cutLen); i < len(arr)-int(cutLen); i++ {
		sum += arr[i]
	}

	return float64(sum) / (float64(len(arr)) - cutLen*2)
}
