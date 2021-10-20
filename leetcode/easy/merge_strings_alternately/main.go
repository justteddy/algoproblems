package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/merge-strings-alternately/
func main() {
	fmt.Println(mergeAlternately("asdasd", "asfjnfaskkm"))
}

func mergeAlternately(word1 string, word2 string) string {
	b := new(strings.Builder)
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			b.WriteByte(word1[i])
		}
		if i < len(word2) {
			b.WriteByte(word2[i])
		}
	}

	return b.String()
}
