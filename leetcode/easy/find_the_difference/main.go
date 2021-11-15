package main

import "fmt"

// https://leetcode.com/problems/find-the-difference/
func main() {
	fmt.Println(findTheDifference("abcd", "abcdd"))
}

func findTheDifference(s string, t string) byte {
	dict := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		cnt := dict[t[i]]
		if cnt == 0 {
			return t[i]
		}
		dict[t[i]]--
	}

	return 0
}
