package main

import "fmt"

// https://leetcode.com/problems/flood-fill/
func main() {
	fmt.Println(floodFill([][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 2, 2, 2},
	}, 0, 0, 1))
}

func floodFill(image [][]int, y int, x int, newColor int) [][]int {
	height := len(image)
	width := len(image[0])

	var floodFillFn func(y, x, prevColor, newColor int)
	floodFillFn = func(y, x, prevColor, newColor int) {
		if x < 0 || x >= width || y < 0 || y >= height {
			return
		}
		if image[y][x] != prevColor || image[y][x] == newColor {
			return
		}

		image[y][x] = newColor
		floodFillFn(y+1, x, prevColor, newColor)
		floodFillFn(y-1, x, prevColor, newColor)
		floodFillFn(y, x+1, prevColor, newColor)
		floodFillFn(y, x-1, prevColor, newColor)
	}

	floodFillFn(y, x, image[y][x], newColor)

	return image
}
