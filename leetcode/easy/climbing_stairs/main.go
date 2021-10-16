package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs/
func main() {
	fmt.Println(climbStairs(20))
}

func climbStairs(n int) int {
	cache := make(map[int]int)
	var ways func(int) int
	ways = func(n int) int {
		if n < 3 {
			return n
		}

		w1, ok := cache[n-1]
		if !ok {
			w1 = ways(n - 1)
			cache[n-1] = w1
		}

		w2, ok := cache[n-2]
		if !ok {
			w2 = ways(n - 2)
			cache[n-2] = w2
		}

		return w1 + w2
	}

	return ways(n)
}
