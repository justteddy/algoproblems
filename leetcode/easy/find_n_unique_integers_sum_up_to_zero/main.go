package main

import "fmt"

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
func main() {
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {
	result := make([]int, 0, n)
	i, to := 1, n/2
	for i <= to {
		result = append(result, []int{i, -i}...)
		i++
	}
	if n&1 == 1 {
		result = append(result, 0)
	}

	return result
}
