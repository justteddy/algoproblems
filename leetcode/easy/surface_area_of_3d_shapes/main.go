package main

import "fmt"

// https://leetcode.com/problems/surface-area-of-3d-shapes/
func main() {
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))
}

func surfaceArea(grid [][]int) int {
	area := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			height := grid[y][x]
			if height == 0 {
				continue
			}

			cubeArea := 2 // bottom and top area

			// top
			if y-1 < 0 {
				cubeArea += height
			} else if grid[y-1][x] < height {
				cubeArea += height - grid[y-1][x]
			}

			// bottom
			if y+1 >= len(grid) {
				cubeArea += height
			} else if grid[y+1][x] < height {
				cubeArea += height - grid[y+1][x]
			}

			// left
			if x-1 < 0 {
				cubeArea += height
			} else if grid[y][x-1] < height {
				cubeArea += height - grid[y][x-1]
			}

			// right
			if x+1 >= len(grid[0]) {
				cubeArea += height
			} else if grid[y][x+1] < height {
				cubeArea += height - grid[y][x+1]
			}

			area += cubeArea
		}
	}

	return area
}
