package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	start := uint8(96)
	result := make([]string, 0)
	i := len(s) - 1
	for i >= 0 {
		if s[i] == '#' {
			num, _ := strconv.Atoi(s[i-2 : i])
			result = append(result, string(start+uint8(num)))
			i -= 3
			continue
		}
		num, _ := strconv.Atoi(string(s[i]))
		result = append(result, string(start+uint8(num)))
		i--
	}

	return strings.Join(reverse(result), "")
}

func reverse(s []string) []string {
	result := make([]string, len(s))
	for i, sym := range s {
		result[len(s)-i-1] = sym
	}

	return result
}
