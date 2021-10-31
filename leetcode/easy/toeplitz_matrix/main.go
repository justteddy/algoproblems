package main

import "fmt"

// https://leetcode.com/problems/toeplitz-matrix/
func main() {
	fmt.Println(isToeplitzMatrix([][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}))
}

func isToeplitzMatrix(matrix [][]int) bool {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if row+1 < len(matrix) && col+1 < len(matrix[0]) {
				if matrix[row][col] != matrix[row+1][col+1] {
					return false
				}
			}
		}
	}

	return true
}
