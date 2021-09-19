package main

import "fmt"

// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
func main() {
	fmt.Println(areOccurrencesEqual("abcabcabc"))
}

func areOccurrencesEqual(s string) bool {
	hm := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		hm[s[i]]++
	}

	prev := 0
	for _, cnt := range hm {
		if prev == 0 {
			prev = cnt
			continue
		}
		if prev != cnt {
			return false
		}
		prev = cnt
	}

	return true
}
