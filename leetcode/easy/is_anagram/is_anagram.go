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

// sum solution
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var alphabet [26]int
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		alphabet[t[i]-'a']--
	}

	for _, a := range alphabet {
		if a != 0 {
			return false
		}
	}

	return true
}

// hash map solution
func isAnagram_(s string, t string) bool {
	dict := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		_, ok := dict[t[i]]
		if !ok {
			return false
		}
		dict[t[i]]--
		if dict[t[i]] == 0 {
			delete(dict, t[i])
		}
	}

	return len(dict) == 0
}
