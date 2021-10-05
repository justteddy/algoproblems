package main

import "fmt"

// https://leetcode.com/problems/shortest-distance-to-a-character/
func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(s string, c byte) []int {
	appearances := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			appearances = append(appearances, i)
		}
	}

	result := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		result = append(result, shortestDistance(appearances, i))
	}

	return result
}

func shortestDistance(appearances []int, i int) int {
	min := 10_000
	for _, appear := range appearances {
		dist := abs(appear - i)
		if dist < min {
			min = dist
		}
	}

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
