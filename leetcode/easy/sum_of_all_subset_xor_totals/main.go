package main

import (
	"fmt"
	"math"
)

/*
	https://leetcode.com/problems/sum-of-all-subset-xor-totals/
*/

func main() {
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
}

var result [][]int

// backtracking
func subsetXORSum(nums []int) int {
	result = make([][]int, 0, 1<<len(nums))
	subsets(nums, []int{}, 0)

	fmt.Println(result)

	sum := 0
	for _, subs := range result {
		xorSum := 0
		for _, v := range subs {
			xorSum ^= v
		}
		sum += xorSum
	}

	return sum
}

func subsets(nums []int, subs []int, index int) {
	result = append(result, subs)

	for i := index; i < len(nums); i++ {
		subs = append(subs, nums[i])
		with := make([]int, len(subs))
		copy(with, subs)

		subsets(nums, with, i+1)
		if len(subs) == 1 {
			subs = []int{}
		} else {
			subs = subs[:len(subs)-1]
		}
	}
}

// solution with OR bitmask
func subsetXORSum_(nums []int) int {
	bits := 0
	for _, n := range nums {
		bits |= n
	}
	return bits * int(math.Pow(2, float64(len(nums)-1)))
}
