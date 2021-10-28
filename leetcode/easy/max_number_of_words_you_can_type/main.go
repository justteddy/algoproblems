package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
func main() {
	fmt.Println(canBeTypedWords("leet code", "e"))
}

func canBeTypedWords(text string, brokenLetters string) int {
	brokenMap := make(map[byte]struct{})
	for i := 0; i < len(brokenLetters); i++ {
		brokenMap[brokenLetters[i]] = struct{}{}
	}

	result, isBroken := 0, false
	for i := 0; i < len(text); i++ {
		if text[i] == 32 { // space char
			if !isBroken {
				result++
			}
			isBroken = false
			continue
		}

		if isBroken { // if word is broken already - continue
			continue
		}
		if _, ok := brokenMap[text[i]]; ok {
			isBroken = true
		}
	}

	// last word
	if !isBroken {
		result++
	}

	return result
}
