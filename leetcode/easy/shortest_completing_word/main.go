package main

import "fmt"

// https://leetcode.com/problems/shortest-completing-word/
func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	dict := createDict(licensePlate)

	result, shortestLen := "", 15
	for _, word := range words {
		m := copyMap(dict)

		for i := 0; i < len(word); i++ {
			if _, ok := m[word[i]]; !ok {
				continue
			}

			m[word[i]]--
			if m[word[i]] == 0 {
				delete(m, word[i])
			}
		}

		if len(m) != 0 {
			continue
		}

		if len(word) < shortestLen {
			shortestLen = len(word)
			result = word
		}
	}

	return result
}

func copyMap(m map[byte]int) map[byte]int {
	result := make(map[byte]int, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}

func createDict(licensePlate string) map[byte]int {
	dict := make(map[byte]int)
	for i := 0; i < len(licensePlate); i++ {
		switch {
		case licensePlate[i] >= 'A' && licensePlate[i] <= 'Z':
			dict[licensePlate[i]+32]++
		case licensePlate[i] >= 'a' && licensePlate[i] <= 'z':
			dict[licensePlate[i]]++
		default:
			continue
		}
	}

	return dict
}
