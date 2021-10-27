package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
func main() {
	fmt.Println(countBalls(1, 100))
}

func countBalls(lowLimit int, highLimit int) int {
	result := 0
	m := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		s := digitSum(i)
		m[s]++
		result = max(m[s], result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func digitSum(num int) int {
	result := 0
	for num != 0 {
		result += num % 10
		num /= 10
	}

	return result
}
