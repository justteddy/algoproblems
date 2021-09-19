package main

import "fmt"

// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			result++
		}
	}

	return result
}
