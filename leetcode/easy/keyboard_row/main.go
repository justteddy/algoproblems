package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/keyboard-row/
func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Keyboard Row.
Memory Usage: 2 MB, less than 94.29% of Go online submissions for Keyboard Row.
*/
func findWords(words []string) []string {
	dict := make(map[string]int, 0)
	dict = makeMap(dict, "qwertyuiop", 1)
	dict = makeMap(dict, "asdfghjkl", 2)
	dict = makeMap(dict, "zxcvbnm", 3)

	// Hello
	result := make([]string, 0)
	for _, word := range words {
		prev := dict[strings.ToLower(string(word[0]))]
		ok := true
		for _, char := range word {
			current := dict[strings.ToLower(string(char))]
			if current != prev {
				ok = false
				break
			}
		}

		if ok {
			result = append(result, word)
		}
	}

	return result
}

func makeMap(dict map[string]int, chars string, n int) map[string]int {
	for _, char := range chars {
		dict[string(char)] = n
	}

	return dict
}
