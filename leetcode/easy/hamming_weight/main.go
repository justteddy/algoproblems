package main

import "fmt"

// https://leetcode.com/problems/number-of-1-bits/
func main() {
	fmt.Println(hammingWeight(192831293))
}

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		if num&1 == 1 {
			result++
		}
		num >>= 1
	}
	return result
}
