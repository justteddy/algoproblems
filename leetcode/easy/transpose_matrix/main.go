package main

import "fmt"

// https://leetcode.com/problems/transpose-matrix/
func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func transpose(matrix [][]int) [][]int {
	result := make([][]int, 0, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		newRow := make([]int, 0, len(matrix))
		for j := 0; j < len(matrix); j++ {
			newRow = append(newRow, matrix[j][i])
		}
		result = append(result, newRow)
	}
	return result
}
