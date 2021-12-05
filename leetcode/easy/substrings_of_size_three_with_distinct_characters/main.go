package main

import "fmt"

// https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/
func main() {
	fmt.Println(countGoodSubstrings("aababcabc"))
}

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}

	result := 0
	start, end := 0, 2
	for end < len(s) {
		if isDistinct(s[start : end+1]) {
			result++
		}
		start++
		end++
	}

	return result
}

func isDistinct(s string) bool {
	d := [26]int{}
	for i := 0; i < len(s); i++ {
		d[s[i]-'a']++
		if d[s[i]-'a'] > 1 {
			return false
		}
	}

	return true
}
