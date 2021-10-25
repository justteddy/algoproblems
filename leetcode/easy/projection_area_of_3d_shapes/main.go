package main

import "fmt"

// https://leetcode.com/problems/projection-area-of-3d-shapes/
func main() {
	fmt.Println(projectionArea([][]int{
		{1, 2},
		{3, 4},
	}))
}

func projectionArea(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		bestRow, bestCol := 0, 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 0 {
				result++
			}
			bestRow = max(grid[i][j], bestRow)
			bestCol = max(grid[j][i], bestCol)
		}
		result += bestRow + bestCol
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
