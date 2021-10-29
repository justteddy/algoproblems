package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/goat-latin/
func main() {
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}

func toGoatLatin(sentence string) string {
	vowels := getVowels()
	words := strings.Split(sentence, " ")

	ma := "ma"
	result := new(strings.Builder)

	for i, word := range words {
		ma += "a"
		if _, ok := vowels[word[0]]; ok {
			result.WriteString(word + ma)
		} else {
			result.Write([]byte(word[1:]))
			result.WriteByte(word[0])
			result.WriteString(ma)
		}
		if i != len(words)-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func getVowels() map[byte]struct{} {
	return map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}
}
