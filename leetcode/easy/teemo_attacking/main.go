package main

import "fmt"

// https://leetcode.com/problems/teemo-attacking/
func main() {
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	result, until := 0, -1
	for _, attack := range timeSeries {
		if attack > until {
			until = attack + duration - 1
			result += duration
			continue
		}
		overflow := attack + duration - 1
		result += overflow - until
		until = overflow
	}
	return result
}
