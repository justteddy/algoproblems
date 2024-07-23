package main

import "fmt"

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
}

func isPrefixOfWord(sentence string, searchWord string) int {
	wrd := 0
	for i, v := range sentence {
		if i == 0 || string(v) == " " {
			wrd++
			idx := i
			if i != 0 {
				idx = i + 1
			}
			if len(sentence[idx:]) < len(searchWord) {
				break
			}
			if sentence[idx:idx+len(searchWord)] == searchWord {
				return wrd
			}
		}
	}

	return -1
}
