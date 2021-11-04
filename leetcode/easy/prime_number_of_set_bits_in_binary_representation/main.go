package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/
func main() {
	fmt.Println(countPrimeSetBits(22, 33))
}

func countPrimeSetBits(left int, right int) int {
	result := 0
	for i := left; i <= right; i++ {
		result += isPrime(bitCount(i))
	}
	return result
}

func isPrime(value int) int {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return 0
		}
	}

	if value > 1 {
		return 1
	}

	return 0
}

func bitCount(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}
