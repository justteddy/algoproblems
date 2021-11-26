package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/relative-sort-array/
func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]int, len(arr2))
	for i, v := range arr2 {
		arr2Map[v] = i
	}

	arr1Map := make(map[int]int, len(arr2))
	needToSort := make([]int, 0, len(arr1)-len(arr2))
	for _, v := range arr1 {
		if _, ok := arr2Map[v]; !ok {
			needToSort = append(needToSort, v)
			continue
		}
		arr1Map[v]++
	}

	sort.Ints(needToSort)

	result := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		result = append(result, fillSlice(v, arr1Map[v])...)
	}

	return append(result, needToSort...)
}

func fillSlice(val, length int) []int {
	result := make([]int, 0, length)
	for i := 1; i <= length; i++ {
		result = append(result, val)
	}

	return result
}
