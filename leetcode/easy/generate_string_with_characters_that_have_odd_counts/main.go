package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
func main() {
	fmt.Println(generateTheString(10))
}

func generateTheString(n int) string {
	if n&1 == 0 {
		return genString(n-1) + "y"
	}
	return genString(n)
}

func genString(n int) string {
	b := new(strings.Builder)
	for i := 1; i <= n; i++ {
		b.WriteString("x")
	}
	return b.String()
}
