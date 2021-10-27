package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/excel-sheet-column-number/
func main() {
	fmt.Println(titleToNumber("FXSHRXW"))
}

func titleToNumber(columnTitle string) int {
	result, position := 0, 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += int(math.Pow(26, float64(position))) * (int(columnTitle[i] - 64))
		position++
	}
	return result
}
