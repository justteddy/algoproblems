package main

import (
	"sort"
	"fmt"
)

// задача 1.
// дан массив чисел, найти минимальное число, которого нет в массиве
// числа неотрицательные, 0 <= len(nums) <= 10^5
// примеры:
// [] -> 0
// [0] -> 1
// [2, 2, 0, 4, 1] -> 3

func main() {
	fmt.Println(findMissingNumber_([]int{2, 2, 0, 4, 1}))
	fmt.Println(findMissingNumber_([]int{}))
	fmt.Println(findMissingNumber_([]int{0}))
}

// O(n) time, O(n) space
func findMissingNumber_(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0] + 1
	}

	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return 0

}

// O(nlogn) time, O(1) space
func findMissingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0] + 1
	}

	sort.Ints(nums)

	var res int
	for i, v := range nums {
		if i == len(nums)-1 {
			res = v + 1
			break
		}

		if v == nums[i+1] {
			continue
		}

		if v != nums[i+1]-1 {
			res = v + 1
			break
		}
	}

	return res
}
