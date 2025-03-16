package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length
func main() {
	fmt.Println(maxVowels("abciiidef", 3)) // 3
}

func maxVowels(s string, k int) int {
	// abciiidef, k=3
	maxVowels, curr := 0, 0
	leftIdx := 0
	for i := 0; i < len(s); i++ {
		if i < k {
			if isVowel(s[i]) {
				maxVowels++
				curr++
			}
			continue
		}

		// move left border of sliding window
		if isVowel(s[leftIdx]) {
			curr--
		}
		leftIdx++

		if isVowel(s[i]) {
			curr++
		}

		if curr > maxVowels {
			maxVowels = curr
		}
	}

	return maxVowels
}

func isVowel(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}

	return false
}
