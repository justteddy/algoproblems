package main

import "fmt"

/*
	You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.
	Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and lowercase letters.
	Return true if a and b are alike. Otherwise, return false.

	Example 1:
	Input: s = "book"
	Output: true
	Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.

	Example 2:
	Input: s = "textbook"
	Output: false
	Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
	Notice that the vowel o is counted twice.

	Example 3:
	Input: s = "MerryChristmas"
	Output: false

	Example 4:
	Input: s = "AbCdEfGh"
	Output: true

	Constraints:
	2 <= s.length <= 1000
	s.length is even.
	s consists of uppercase and lowercase letters.
*/
func main() {
	fmt.Println(halvesAreAlike("book"))
}

func halvesAreAlike(s string) bool {
	vowels := map[byte]struct{}{
		97:  struct{}{},
		101: struct{}{},
		105: struct{}{},
		111: struct{}{},
		117: struct{}{},
		65:  struct{}{},
		69:  struct{}{},
		73:  struct{}{},
		79:  struct{}{},
		85:  struct{}{},
	}

	count := 0
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if _, ok := vowels[s[i]]; ok {
			count++
		}
		if _, ok := vowels[s[j]]; ok {
			count--
		}
	}

	return count == 0
}
