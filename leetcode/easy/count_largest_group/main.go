package main

import "fmt"

// https://leetcode.com/problems/count-largest-group/
func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	result, max := 0, 0
	m := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := sumByDigits(i)
		m[sum]++
		if m[sum] > max {
			max = m[sum]
			result = 1
		} else if m[sum] == max {
			result++
		}
	}

	return result
}

func sumByDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}
