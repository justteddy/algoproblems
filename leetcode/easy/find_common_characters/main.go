package main

import "fmt"

// https://leetcode.com/problems/find-common-characters/
func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(words []string) []string {
	dict := make(map[uint8]int)
	for i := 0; i < len(words[0]); i++ {
		dict[words[0][i]]++
	}

	for i := 1; i < len(words); i++ {
		wordDict := make(map[uint8]int)
		for j := 0; j < len(words[i]); j++ {
			if _, ok := dict[words[i][j]]; ok {
				wordDict[words[i][j]]++
			}
		}

		for key := range wordDict {
			wordDict[key] = min(wordDict[key], dict[key])
		}
		dict = wordDict
	}

	result := make([]string, 0)
	for key, cnt := range dict {
		for i := 1; i <= cnt; i++ {
			result = append(result, string(key))
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
