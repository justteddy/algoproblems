package main

import "fmt"

// https://leetcode.com/problems/01-matrix/
func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))
}

func updateMatrix(mat [][]int) [][]int {
	queue := make([][2]int, 0)
	result := make([][]int, 0, len(mat))
	for i, rows := range mat {
		resultRow := make([]int, 0, len(mat[0]))
		for j := range rows {
			if mat[i][j] == 0 {
				resultRow = append(resultRow, 0)
				queue = append(queue, [2]int{i, j})
			} else {
				resultRow = append(resultRow, 10_000)
			}
		}
		result = append(result, resultRow)
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			y, x := curr[0]+dir[0], curr[1]+dir[1]
			if y < 0 || x < 0 || y >= len(result) || x >= len(result[0]) {
				continue
			}

			if result[y][x] > result[curr[0]][curr[1]]+1 {
				result[y][x] = result[curr[0]][curr[1]] + 1
				queue = append(queue, [2]int{y, x})
			}
		}

	}

	return result
}
