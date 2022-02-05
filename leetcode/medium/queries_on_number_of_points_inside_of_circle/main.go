package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/
func main() {
	fmt.Println(countPoints([][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}}))
}

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, 0, len(queries))
	for _, query := range queries {
		inside := 0
		for _, point := range points {
			xDist := math.Pow(float64(point[0]-query[0]), 2)
			yDist := math.Pow(float64(point[1]-query[1]), 2)
			distance := math.Sqrt(xDist + yDist)
			if distance <= float64(query[2]) {
				inside++
			}
		}
		result = append(result, inside)
	}

	return result
}
