package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCasePermutation("A1b2"))
}

func letterCasePermutation(s string) []string {
	result := make([]string, 0)
	permutations(0, strings.ToLower(s), &result)
	return result
}

func permutations(index int, s string, result *[]string) {
	if index == len(s) {
		*result = append(*result, s)
		return
	}

	if s[index] >= 'a' && s[index] <= 'z' {
		bytes := []byte(s)
		bytes[index] -= 32
		permutations(index+1, string(bytes), result)
	}

	permutations(index+1, s, result)
}
