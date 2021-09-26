package main

import "fmt"

// https://leetcode.com/problems/permutation-in-string/
func main() {
	fmt.Println(checkInclusion("ab", "eidb"))
}

// sliding window
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Map, s2Map [26]int
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		if isMatch(s1Map, s2Map) {
			return true
		}
		if i+len(s1) == len(s2) {
			return false
		}
		s2Map[s2[i+len(s1)]-'a']++
		s2Map[s2[i]-'a']--
	}
	return false
}

// array solution
func checkInclusion_(s1 string, s2 string) bool {
	s1Map := toArray(s1)
	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Map := toArray(s2[i : i+len(s1)])
		if isMatch(s1Map, s2Map) {
			return true
		}
	}
	return false
}

func toArray(s string) [26]int {
	var arr [26]int
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	return arr
}

func isMatch(s1, s2 [26]int) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
