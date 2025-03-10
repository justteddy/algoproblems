package main

import "fmt"

// https://leetcode.com/problems/reverse-vowels-of-a-string
func main() {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	x := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		if !isVowel(s[i]) {
			i++
			continue
		}

		if !isVowel(s[j]) {
			j--
			continue
		}

		x[i], x[j] = x[j], x[i]
		i++
		j--
	}

	return string(x)
}

func isVowel(l byte) bool {
	if l == 'a' || l == 'A' ||
		l == 'e' || l == 'E' ||
		l == 'i' || l == 'I' ||
		l == 'o' || l == 'O' ||
		l == 'u' || l == 'U' {
		return true
	}

	return false
}
