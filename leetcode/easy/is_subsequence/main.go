package main

// https://leetcode.com/problems/is-subsequence
func isSubsequence(s string, t string) bool {
	if s == "" && t == "" {
		return true
	}

	if t == "" {
		return false
	}

	if s == "" {
		return true
	}

	j := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j >= len(s) {
				return true
			}
		}
	}

	return false
}
