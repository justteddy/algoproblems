package main

import "fmt"

// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
func main() {
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	targetSum, pow := 0, 1
	for i := len(targetWord) - 1; i >= 0; i-- {
		targetSum += int(targetWord[i]-'a') * pow
		pow *= 10
	}

	for _, word := range []string{firstWord, secondWord} {
		pow = 1
		for i := len(word) - 1; i >= 0; i-- {
			targetSum -= int(word[i]-'a') * pow
			pow *= 10
			if targetSum < 0 {
				return false
			}
		}
	}

	return targetSum == 0
}
