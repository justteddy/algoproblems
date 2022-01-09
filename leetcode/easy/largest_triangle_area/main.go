package main

import (
	"fmt"
)

// https://leetcode.com/problems/largest-triangle-area/
func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}

func largestTriangleArea(points [][]int) float64 {
	var max float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				s := area(points[i], points[j], points[k])
				if s > max {
					max = s
				}
			}
		}
	}
	return max
}

func area(p1, p2, p3 []int) float64 {
	return abs(float64(p1[0]*(p2[1]-p3[1])+p2[0]*(p3[1]-p1[1])+p3[0]*(p1[1]-p2[1])) / 2)
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
