package main

import "fmt"

// https://leetcode.com/problems/first-unique-character-in-a-string/
func main() {
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
	dict := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			return i
		}
	}

	return -1
}
