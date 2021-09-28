package main

import "fmt"

// https://leetcode.com/problems/pascals-triangle/
func main() {
	fmt.Println(generate(8))
}

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		if i < 3 {
			for j := range row {
				row[j] = 1
			}
		} else {
			for j := range row {
				if j == 0 || j == len(row)-1 {
					row[j] = 1
				} else {
					row[j] = result[i-2][j-1] + result[i-2][j]
				}
			}
		}
		result = append(result, row)
	}
	return result
}
