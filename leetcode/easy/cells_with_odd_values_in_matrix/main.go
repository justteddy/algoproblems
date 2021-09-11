package main

import "fmt"

/*
	There is an m x n matrix that is initialized to all 0's. There is also a 2D array indices where each indices[i] = [ri, ci] represents a 0-indexed location to perform some increment operations on the matrix.
	For each location indices[i], do both of the following:
	Increment all the cells on row ri.
	Increment all the cells on column ci.
	Given m, n, and indices, return the number of odd-valued cells in the matrix after applying the increment to all locations in indices.

	Example 1:
	Input: m = 2, n = 3, indices = [[0,1],[1,1]]
	Output: 6
	Explanation: Initial matrix = [[0,0,0],[0,0,0]].
	After applying first increment it becomes [[1,2,1],[0,1,0]].
	The final matrix is [[1,3,1],[1,3,1]], which contains 6 odd numbers.

	Example 2:
	Input: m = 2, n = 2, indices = [[1,1],[0,0]]
	Output: 0
	Explanation: Final matrix = [[2,2],[2,2]]. There are no odd numbers in the final matrix.

	Constraints:
	1 <= m, n <= 50
	1 <= indices.length <= 100
	0 <= ri < m
	0 <= ci < n
*/
func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
}

func oddCells(m int, n int, indices [][]int) int {
	ans := 0
	matrix := buildMatrix(m, n)
	for _, index := range indices {
		for i := range matrix[index[0]] {
			matrix[index[0]][i]++
			if matrix[index[0]][i]&1 == 1 {
				ans++
			} else {
				ans--
			}
		}
		for _, row := range matrix {
			row[index[1]]++
			if row[index[1]]&1 == 1 {
				ans++
			} else {
				ans--
			}
		}
	}

	return ans
}

func buildMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}
