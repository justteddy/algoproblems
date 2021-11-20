package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/uncommon-words-from-two-sentences/
func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func uncommonFromSentences(s1 string, s2 string) []string {
	result := make([]string, 0)
	dict := make(map[string]int)

	for _, w := range strings.Split(s1, " ") {
		dict[w]++
	}

	for _, w := range strings.Split(s2, " ") {
		dict[w]++
	}

	for w, cnt := range dict {
		if cnt == 1 {
			result = append(result, w)
		}
	}

	return result
}
