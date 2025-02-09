package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	digits := make([]int, 0)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	res := 0
	for i := 0; i < len(digits); i++ {
		res += digits[i] * pow10(len(digits)-i-1)
	}

	if res < -math.MaxInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

func pow10(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= 10
	}

	return res
}
