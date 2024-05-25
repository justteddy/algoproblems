package main

import "fmt"

func main() {
	fmt.Println(countLetters("aaaba"))
	fmt.Println(countLetters("aaaaaaaaaa"))
}

func countLetters(s string) int {
	res := 0
	for i, j := 0, 1; i < len(s)-1; {
		if s[i] == s[j] {
			res++
			j++
			if j == len(s) {
				i++
				j = i + 1
			}
		} else {
			i++
			j = i + 1
		}
	}

	return res + len(s)
}
