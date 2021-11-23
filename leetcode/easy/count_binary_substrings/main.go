package main

import "fmt"

// https://leetcode.com/problems/count-binary-substrings/
func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}

func countBinarySubstrings(s string) int {
	if len(s) == 1 {
		return 0
	}

	groups := make([]int, 0)
	prev, cnt := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			cnt++
			continue
		}
		groups = append(groups, cnt)
		prev = s[i]
		cnt = 1
	}

	if cnt > 0 {
		groups = append(groups, cnt)
	}

	result := 0
	for i := 0; i < len(groups)-1; i++ {
		result += min(groups[i], groups[i+1])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
