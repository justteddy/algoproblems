package main

import "fmt"

// https://leetcode.com/problems/happy-number/
func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	cache := make(map[int]int)
	digitsSquareFn := func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}

		result := 0
		for n != 0 {
			digit := n % 10
			result += digit * digit
			n /= 10
		}

		cache[n] = result
		return result
	}

	prev := n
	previousSums := make(map[int]struct{})
	for {
		digitSquareSum := digitsSquareFn(prev)
		if _, ok := previousSums[digitSquareSum]; ok {
			return false
		}

		if digitSquareSum == 1 {
			return true
		}
		previousSums[digitSquareSum] = struct{}{}
		prev = digitSquareSum
	}
}
