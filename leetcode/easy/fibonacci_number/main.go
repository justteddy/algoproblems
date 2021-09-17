package main

import "fmt"

// https://leetcode.com/problems/fibonacci-number/
func main() {
	fmt.Println(fib(6))
}

var hm = map[int]int{}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	if memo, ok := hm[n]; ok {
		return memo
	}

	hm[n] = fib(n-1) + fib(n-2)

	return hm[n]
}
