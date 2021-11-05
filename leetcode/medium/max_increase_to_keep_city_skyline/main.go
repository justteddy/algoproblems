package main

import "fmt"

// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMaxs, colMaxs := make([]int, 0, len(grid)), make([]int, 0, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		bestRow, bestCol := 0, 0
		for j := 0; j < len(grid); j++ {
			bestRow = max(grid[i][j], bestRow)
			bestCol = max(grid[j][i], bestCol)
		}
		rowMaxs = append(rowMaxs, bestRow)
		colMaxs = append(colMaxs, bestCol)
	}

	result := 0
	for i, rows := range grid {
		for j, col := range rows {
			if rowMaxs[i] > col && colMaxs[j] > col {
				result += min(rowMaxs[i]-col, colMaxs[j]-col)
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
