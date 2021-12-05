package main

import "fmt"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}

func average(salary []int) float64 {
	min, max := 1_000_000, 0
	for _, sal := range salary {
		if sal > max {
			max = sal
		}
		if sal < min {
			min = sal
		}
	}

	var sum float64
	for _, sal := range salary {
		if sal == min || sal == max {
			continue
		}
		sum += float64(sal)
	}

	return sum / (float64(len(salary) - 2))
}
