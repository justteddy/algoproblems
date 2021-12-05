package main

import "fmt"

// https://leetcode.com/problems/palindrome-number/
func main() {
	fmt.Println(isPalindrome(100))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	ints := make([]int, 0)
	for x != 0 {
		ints = append(ints, x%10)
		x /= 10
	}

	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		if ints[i] != ints[j] {
			return false
		}
	}
	return true
}
