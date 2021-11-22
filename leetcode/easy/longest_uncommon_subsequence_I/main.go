package main

import "fmt"

// https://leetcode.com/problems/longest-uncommon-subsequence-i/
func main() {
	fmt.Println(findLUSlength("abc", "asdsbsdasx"))
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	return max(len(a), len(b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
