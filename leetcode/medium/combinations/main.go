package main

import "fmt"

// https://leetcode.com/problems/combinations/
func main() {
	fmt.Println(combine(5, 3))
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	comb(n, k, 1, []int{}, &result)
	return result
}

func comb(n, k, index int, data []int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(data))
		copy(tmp, data)
		*result = append(*result, tmp)
		return
	}

	for i := index; i <= n-k+1; i++ {
		data = append(data, i)
		comb(n, k-1, i+1, data, result)
		data = data[:len(data)-1]
	}
}
