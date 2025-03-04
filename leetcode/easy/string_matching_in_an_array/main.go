package main

import (
	"strings"
	"fmt"
)

// https://leetcode.com/problems/string-matching-in-an-array/
func main() {
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}

// brute force, with usage of stdlib
func stringMatching(words []string) []string {
	m := make(map[string]bool)
	for _, w := range words {
		for _, w2 := range words {
			if w == w2 {
				continue
			}
			if strings.Contains(w, w2) {
				m[w2] = true
			}
		}
	}

	res := make([]string, 0, len(m))
	for k := range m {
		res = append(res, k)
	}

	return res
}
