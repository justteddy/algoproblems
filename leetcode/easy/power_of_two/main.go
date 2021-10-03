package main

import "fmt"

// https://leetcode.com/problems/power-of-two/
func main() {
	fmt.Println(isPowerOfTwo(64))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n&1 != 0 {
			return false
		}
		n = n >> 1
	}
	return true
}
