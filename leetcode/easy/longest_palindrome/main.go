package main

import "fmt"

// https://leetcode.com/problems/longest-palindrome/
func main() {
	fmt.Println(longestPalindrome("bananas"))
}

/*
 abccccdd -> 7 dccaccd
 bananas  -> 5 naban
 cccaaa   -> 5 cacac
 cccaaa   -> 5 cacac
 cccaaaa  -> caacaac
 ccc      -> 3 ccc
 c		  -> 1 c
 cc       -> 2 cc
*/

// naive solution
func longestPalindrome(s string) int {
	res := 0
	m := make(map[string]int)
	for _, sym := range s {
		v := string(sym)
		if _, ok := m[v]; !ok {
			m[v] = 1
			continue
		}
		m[v]++
	}

	if len(m) == 1 {
		return len(s)
	}

	single := false
	for _, v := range m {
		if v == 1 && !single {
			single = true
			res++
			continue
		}

		if v%2 == 1 {
			res += v - 1
			if !single {
				res++
				single = true
				continue
			}
		}

		if v%2 == 0 {
			res += v
		}
	}

	return res
}

// optimal
func longestPalindrome2(s string) int {
	var res int
	dict := make(map[int32]int)
	for _, c := range s {
		dict[c]++
		num := dict[c]
		if num == 2 {
			res = res + 2
			delete(dict, c)
		}
	}

	if len(dict) != 0 {
		res++
	}

	return res
}
