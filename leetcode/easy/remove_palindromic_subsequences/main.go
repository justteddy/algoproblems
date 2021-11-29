package main

import "fmt"

// https://leetcode.com/problems/remove-palindromic-subsequences/
func main() {
	fmt.Println(removePalindromeSub("ababa"))
}

func removePalindromeSub(s string) int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		return 2
	}

	return 1
}
