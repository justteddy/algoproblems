package main

import "fmt"

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/
func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}))
}

func canBeEqual(target []int, arr []int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(target); i++ {
		dict[target[i]]++
		dict[arr[i]]--
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}
