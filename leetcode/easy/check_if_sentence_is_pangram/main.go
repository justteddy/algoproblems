package main

import "fmt"

/*
	A pangram is a sentence where every letter of the English alphabet appears at least once.

	Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

	Example 1:
	Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
	Output: true
	Explanation: sentence contains at least one of every letter of the English alphabet.

	Example 2:
	Input: sentence = "leetcode"
	Output: false

	Constraints:
	1 <= sentence.length <= 1000
	sentence consists of lowercase English letters.
*/
func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}

func checkIfPangram(sentence string) bool {
	hm := make(map[uint8]struct{})
	for i := 97; i <= 122; i++ {
		hm[uint8(i)] = struct{}{}
	}

	for i := 0; i < len(sentence); i++ {
		delete(hm, sentence[i])
		if len(hm) == 0 {
			return true
		}
	}

	return len(hm) == 0
}
