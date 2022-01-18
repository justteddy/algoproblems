package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reformat-phone-number/
func main() {
	fmt.Println(reformatNumber("1-23-45 6"))
}

func reformatNumber(number string) string {
	number = cleanNumber(number)

	result := make([]string, 0)
	i := 0
OUTER:
	for {
		switch len(number[i:]) {
		case 4:
			result = append(result, number[i:i+2])
			result = append(result, number[i+2:])
			break OUTER
		case 2, 3:
			result = append(result, number[i:])
			break OUTER
		default:
			result = append(result, number[i:i+3])
			i += 3
		}
	}

	return strings.Join(result, "-")
}

func cleanNumber(number string) string {
	result := make([]string, 0)
	for _, v := range number {
		if v >= '0' && v <= '9' {
			result = append(result, string(v))
		}
	}
	return strings.Join(result, "")
}
