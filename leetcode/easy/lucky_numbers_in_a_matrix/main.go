package main

import "fmt"

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
func main() {
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
}

func luckyNumbers(matrix [][]int) []int {
	dict := make(map[int]int)

	for i := 0; i < len(matrix); i++ {
		minInRow := 100_000
		for j := 0; j < len(matrix[0]); j++ {
			minInRow = min(matrix[i][j], minInRow)
		}
		dict[minInRow]++
	}

	for i := 0; i < len(matrix[0]); i++ {
		maxInCol := 0
		for j := 0; j < len(matrix); j++ {
			maxInCol = max(matrix[j][i], maxInCol)
		}
		dict[maxInCol]++
	}

	result := make([]int, 0)
	for num, count := range dict {
		if count == 2 {
			result = append(result, num)
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
