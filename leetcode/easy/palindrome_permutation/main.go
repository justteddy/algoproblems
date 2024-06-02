package main

import "fmt"

func main() {
	fmt.Println(canPermutePalindrome("code"))
	fmt.Println(canPermutePalindrome("aab"))
	fmt.Println(canPermutePalindrome("baab"))
	fmt.Println(canPermutePalindrome("carerac"))
}

func canPermutePalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	oddCnt := 0
	for _, cnt := range m {
		if cnt&1 == 0 {
			continue
		}
		oddCnt++
		if oddCnt > 1 {
			return false
		}
	}

	return true
}
