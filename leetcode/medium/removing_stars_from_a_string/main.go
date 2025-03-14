package main

import (
	"strings"
	"fmt"
)

// https://leetcode.com/problems/removing-stars-from-a-string
func main() {
	fmt.Println(removeStars("leet**cod*e"))
}

func removeStars(s string) string {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v == '*' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, v)
	}

	var res strings.Builder
	for _, v := range stack {
		res.WriteRune(v)
	}

	return res.String()
}
