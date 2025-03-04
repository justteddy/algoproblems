package main

// https://leetcode.com/problems/reverse-only-letters/
func reverseOnlyLetters(s string) string {
	x := []byte(s)
	for i, j := 0, len(x)-1; i < j; {
		if !isLetter(x[i]) {
			i++
			continue
		}

		if !isLetter(x[j]) {
			j--
			continue
		}

		x[i], x[j] = x[j], x[i]
		i++
		j--
	}

	return string(x)
}

func isLetter(s byte) bool {
	if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) {
		return true
	}

	return false
}
