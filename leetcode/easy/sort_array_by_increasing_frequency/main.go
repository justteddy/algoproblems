package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/sort-array-by-increasing-frequency/
func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
}

// 2 step map solution
func frequencySort(nums []int) []int {
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}

	freqMap := make(map[int][]int)
	for num, freq := range dict {
		if _, ok := freqMap[freq]; !ok {
			freqMap[freq] = []int{num}
			continue
		}
		freqMap[freq] = append(freqMap[freq], num)
	}

	result := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if _, ok := freqMap[i]; !ok {
			continue
		}

		if len(freqMap[i]) == 1 {
			result = append(result, repeat(freqMap[i][0], i)...)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(freqMap[i])))
			for _, val := range freqMap[i] {
				result = append(result, repeat(val, i)...)
			}
		}
	}

	return result
}

func repeat(x int, count int) []int {
	result := make([]int, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, x)
	}

	return result
}
