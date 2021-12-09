package main

import "fmt"

// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 1}, {3, 3}, {5, 5}}))
}

func nearestValidPoint(x int, y int, points [][]int) int {
	idx, smallest := 0, 9_999
	found := false
	for i, point := range points {
		if point[0] == x || point[1] == y {
			found = true
			manhattanDistance := abs(point[0]-x) + abs(point[1]-y)
			if (abs(point[0]-x) + abs(point[1]-y)) < smallest {
				idx = i
				smallest = manhattanDistance
			}
		}
	}

	if !found {
		return -1
	}

	return idx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
