package main

import "fmt"

func main() {
	fmt.Print(flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}))
}

func flipAndInvertImage(image [][]int) [][]int {
	for _, block := range image {
		for i, j := 0, len(block)-1; i <= j; i, j = i+1, j-1 {
			block[i], block[j] = block[j]^1, block[i]^1
		}
	}

	return image
}
