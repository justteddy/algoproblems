package main

import "fmt"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if (i + len(needle)) > len(haystack) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
