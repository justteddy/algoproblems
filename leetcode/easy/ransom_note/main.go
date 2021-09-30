package main

import "fmt"

// https://leetcode.com/problems/ransom-note/
func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	dict := make(map[uint8]int)
	for i := 0; i < len(magazine); i++ {
		dict[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		cnt, ok := dict[ransomNote[i]]
		if !ok || cnt == 0 {
			return false
		}
		dict[ransomNote[i]]--
	}

	return true
}
