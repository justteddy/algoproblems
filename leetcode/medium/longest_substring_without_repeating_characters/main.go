package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// O(n) solution
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	start, result := 0, 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok && m[s[i]] >= start {
			start = m[s[i]] + 1
		}
		result = max(result, i-start+1)
		m[s[i]] = i
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// easy solution O(n*Logn)
func lengthOfLongestSubstring_(s string) int {
	max := 0
	m := make(map[uint8]struct{})
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				m = make(map[uint8]struct{})
				break
			}
			m[s[j]] = struct{}{}
			if len(m) > max {
				max = len(m)
			}
		}
	}
	return max
}
