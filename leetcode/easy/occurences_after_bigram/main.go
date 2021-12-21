package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/occurrences-after-bigram/
func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
}

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	if len(text) <= 2 {
		return nil
	}

	result := make([]string, 0)
	for i, j := 0, 1; j < len(words)-1; i, j = i+1, j+1 {
		if words[i] == first && words[j] == second {
			result = append(result, words[j+1])
		}
	}

	return result
}
