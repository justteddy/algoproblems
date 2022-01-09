package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/count-square-sum-triples/
func main() {
	fmt.Println(countTriples(10))
}

func countTriples(n int) int {
	result := 0
	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			c := math.Sqrt(float64(i*i + j*j))
			x := float64(int(c))
			if c-x == 0 && int(c) <= n {
				result += 2
			}
		}
	}

	return result
}
