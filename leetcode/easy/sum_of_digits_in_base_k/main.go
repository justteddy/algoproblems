package main

import "fmt"

// https://leetcode.com/problems/sum-of-digits-in-base-k/
func main() {
	fmt.Println(sumBase(3450, 13))
}

func sumBase(n int, k int) int {
	result := 0
	for n > 0 {
		result += n % k
		n = n / k
	}
	return result
}
