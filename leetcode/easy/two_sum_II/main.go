package main

import "fmt"

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	dict := make(map[int]int, len(numbers))
	for i, num := range numbers {
		if _, ok := dict[target-num]; ok {
			if i < dict[target-num] {
				return []int{i + 1, dict[target-num] + 1}
			}
			return []int{dict[target-num] + 1, i + 1}
		}
		dict[num] = i
	}
	return nil
}
