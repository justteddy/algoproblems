package main

func main() {

}

// https://leetcode.com/problems/subrectangle-queries/
type SubrectangleQueries struct {
	matrix [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		matrix: rectangle,
	}
}

func (s *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			s.matrix[i][j] = newValue
		}
	}
}

func (s *SubrectangleQueries) GetValue(row int, col int) int {
	return s.matrix[row][col]
}
