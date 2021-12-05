package main

import "fmt"

// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	dict := buildDict(chars)
	result := 0
	for _, word := range words {
		dictCopy := dict
		good := true
		for i := 0; i < len(word); i++ {
			if dictCopy[word[i]-'a'] == 0 {
				good = false
				break
			}
			dictCopy[word[i]-'a']--
		}
		if good {
			result += len(word)
		}
	}

	return result
}

func buildDict(chars string) [26]int {
	d := [26]int{}
	for i := 0; i < len(chars); i++ {
		d[chars[i]-'a']++
	}
	return d
}
