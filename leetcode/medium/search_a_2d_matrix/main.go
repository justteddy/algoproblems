package main

import "fmt"

// https://leetcode.com/problems/search-a-2d-matrix/
func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 11))
}

func searchMatrix(matrix [][]int, target int) bool {
	// search row
	lo, hi := 0, len(matrix)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if matrix[mid][0] == target {
			return true
		}
		// precondition is neccessary
		if target > matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			lo, hi = mid, mid
			break
		}
		if matrix[mid][0] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	// search cell
	row := lo
	lo, hi = 0, len(matrix[0])-1
	for lo < hi {
		mid := (lo + hi) / 2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return matrix[row][lo] == target
}
