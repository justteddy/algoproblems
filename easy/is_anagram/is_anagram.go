package is_anagram

/*
	Given two strings s and t , write a function to determine if t is an anagram of s.

	Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

	Example 2:
	Input: s = "rat", t = "car"
	Output: false
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[uint8]int)
	for i := 0; i <= len(s)-1; i++ {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]] = 1
			continue
		}
		dict[s[i]]++
	}

	for i := 0; i <= len(t)-1; i++ {
		if _, ok := dict[t[i]]; !ok {
			return false
		}
		dict[t[i]]--
		if dict[t[i]] < 0 {
			return false
		}
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}
