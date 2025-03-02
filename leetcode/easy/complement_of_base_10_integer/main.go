package main

import (
	"math/bits"
	"fmt"
)

func main() {
	fmt.Println(bitwiseComplement(5))
}

// https://leetcode.com/problems/complement-of-base-10-integer/
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	mask := (1 << bits.Len(uint(n))) - 1
	fmt.Printf("mask: %b\n", mask)
	fmt.Printf("n: %b\n", n)

	return int(mask ^ n)
}
