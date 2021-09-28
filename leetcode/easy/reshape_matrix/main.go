package main

import "fmt"

// https://leetcode.com/problems/reshape-the-matrix/
func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	result := make([][]int, 0, r)
	tmpCols := make([]int, 0, c)
	for _, rows := range mat {
		for _, col := range rows {
			tmpCols = append(tmpCols, col)
			if len(tmpCols) == c {
				result = append(result, tmpCols)
				tmpCols = make([]int, 0, c)
			}
		}
	}

	return result
}
