package main

import "fmt"

// https://leetcode.com/problems/triangle/
func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

// recursion solution, too slow
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	cache := make(map[string]int)

	var sum func(level, idx int) int
	sum = func(level, idx int) int {
		if level == len(triangle)-1 {
			return triangle[level][idx]
		}

		s1, ok := cache[hash(level+1, idx)]
		if !ok {
			s1 = sum(level+1, idx)
		}

		cache[hash(level+1, idx)] = s1

		s2, ok := cache[hash(level+1, idx+1)]
		if !ok {
			s2 = sum(level+1, idx+1)
		}

		cache[hash(level+1, idx+1)] = s2

		return min(
			triangle[level][idx]+s1,
			triangle[level][idx]+s2,
		)
	}

	return min(triangle[0][0]+sum(1, 0), triangle[0][0]+sum(1, 1))
}

func hash(level, idx int) string {
	return fmt.Sprintf("%d.%d", level, idx)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
