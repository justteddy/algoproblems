package main

import (
	"fmt"
	"strings"
)

/*
	Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

	Example 1:
	Input: s = "Let's take LeetCode contest"
	Output: "s'teL ekat edoCteeL tsetnoc"

	Example 2:
	Input: s = "God Ding"
	Output: "doG gniD"

	Constraints:
	1 <= s.length <= 5 * 104
	s contains printable ASCII characters.
	s does not contain any leading or trailing spaces.
	There is at least one word in s.
	All the words in s are separated by a single space.
*/

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	b := new(strings.Builder)
	words := strings.Split(s, " ")
	for i, word := range words {
		reversed := []byte(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = word[j], word[i]
		}
		b.Write(reversed)
		if i != len(words)-1 {
			b.WriteString(" ")
		}
	}

	return b.String()
}
