package main

import "fmt"

// https://leetcode.com/problems/valid-palindrome/
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	filtered := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			filtered = append(filtered, s[i]+32) // convert to lower case
		}

		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			filtered = append(filtered, s[i])
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i++ {
		if filtered[i] != filtered[j] {
			return false
		}
		j--
	}

	return true
}
