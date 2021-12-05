package main

import "fmt"

// https://leetcode.com/problems/binary-number-with-alternating-bits/
func main() {
	fmt.Println(hasAlternatingBits(5))
}

func hasAlternatingBits(n int) bool {
	prev := n & 1
	for n != 0 {
		n >>= 1
		if n&1 == prev {
			return false
		}
		prev = n & 1
	}

	return true
}
