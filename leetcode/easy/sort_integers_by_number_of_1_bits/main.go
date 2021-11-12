package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
func main() {
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iBitCnt := bitCount(arr[i])
		jBitCnt := bitCount(arr[j])
		if iBitCnt == jBitCnt {
			return arr[i] < arr[j]
		}
		return iBitCnt < jBitCnt
	})

	return arr
}

func bitCount(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}
