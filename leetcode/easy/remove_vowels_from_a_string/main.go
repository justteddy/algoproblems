package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders"))
}

func removeVowels(s string) string {
	// aeiou
	excl := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	b := strings.Builder{}
	for _, l := range s {
		if _, ok := excl[l]; ok {
			continue
		}
		b.WriteRune(l)
	}

	return b.String()
}
