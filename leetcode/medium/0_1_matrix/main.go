package main

import "fmt"

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))
}

func updateMatrix(mat [][]int) [][]int {
	seen := make(map[string]struct{})
	var distanceToZero func(y, x int) int
	distanceToZero = func(y, x int) int {
		if x < 0 || x >= len(mat[0]) || y < 0 || y >= len(mat) {
			return 10_000
		}

		if mat[y][x] == 0 {
			return 0
		}

		if _, ok := seen[hash(y, x)]; ok {
			return 1
		}

		seen[hash(y, x)] = struct{}{}
		return 1 + min(
			distanceToZero(y+1, x),
			distanceToZero(y-1, x),
			distanceToZero(y, x+1),
			distanceToZero(y, x-1),
		)
	}

	result := make([][]int, 0, len(mat))
	for i, rows := range mat {
		resultRow := make([]int, 0, len(mat[0]))
		for j := range rows {
			seen = make(map[string]struct{})
			if mat[i][j] == 0 {
				resultRow = append(resultRow, 0)
				continue
			}
			resultRow = append(resultRow, distanceToZero(i, j))
		}
		result = append(result, resultRow)
	}

	return result
}

func hash(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}

func min(vals ...int) int {
	min := 10_000
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
