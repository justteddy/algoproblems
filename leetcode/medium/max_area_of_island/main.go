package main

import "fmt"

// https://leetcode.com/problems/max-area-of-island/
func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	}))
}

func maxAreaOfIsland(grid [][]int) int {
	seen := make(map[string]struct{})
	var checkArea func(y, x int) int
	checkArea = func(y, x int) int {
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || grid[y][x] != 1 {
			return 0
		}
		if _, ok := seen[hash(y, x)]; ok {
			return 0
		}

		seen[hash(y, x)] = struct{}{}
		return 1 + checkArea(y+1, x) +
			checkArea(y-1, x) +
			checkArea(y, x+1) +
			checkArea(y, x-1)
	}

	max := 0
	for i, rows := range grid {
		for j, col := range rows {
			if col == 0 {
				continue
			}
			if _, ok := seen[hash(i, j)]; ok {
				continue
			}
			area := checkArea(i, j)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func hash(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}
