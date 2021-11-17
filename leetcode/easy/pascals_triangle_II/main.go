package main

import "fmt"

// https://leetcode.com/problems/pascals-triangle-ii/
func main() {
	fmt.Println(getRow(10))
}

func getRow(rowIndex int) []int {
	buffer := make([]int, 0)
	for i := 0; i <= rowIndex; i++ {
		row := make([]int, i+1)
		if i < 2 {
			for j := range row {
				row[j] = 1
			}
		} else {
			for j := range row {
				if j == 0 || j == len(row)-1 {
					row[j] = 1
				} else {
					row[j] = buffer[j-1] + buffer[j]
				}
			}
		}
		buffer = row
	}

	return buffer
}
