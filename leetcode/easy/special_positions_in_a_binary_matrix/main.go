package main

// https://leetcode.com/problems/special-positions-in-a-binary-matrix/description/
func numSpecial(mat [][]int) int {
	rowsMap := make(map[int]int)
	colsMap := make(map[int]int)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				continue
			}
			rowsMap[i]++
			colsMap[j]++
		}
	}

	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				continue
			}
			if rowsMap[i] == 1 && colsMap[j] == 1 {
				res++
			}
		}
	}

	return res
}
