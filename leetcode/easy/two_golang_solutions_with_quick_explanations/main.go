package main

// https://leetcode.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	allUp, allLow, firstUpper := true, true, true

	for i := 0; i < len(word); i++ {
		if !allUp && !allLow && !firstUpper {
			return false
		}

		if word[i] >= 'A' && word[i] <= 'Z' {
			allLow = false
			if i != 0 {
				firstUpper = false
			}
		} else {
			if i == 0 {
				firstUpper = false
			}
			allUp = false
		}
	}

	return firstUpper || allUp || allLow
}
