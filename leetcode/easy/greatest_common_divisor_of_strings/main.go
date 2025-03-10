package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))

}

// https://leetcode.com/problems/greatest-common-divisor-of-strings
func gcdOfStrings(str1 string, str2 string) string {
	var maxDivisor string
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			return ""
		}

		divisor := str1[:i+1]
		if !(checkWithDivisor(str1, divisor) && checkWithDivisor(str2, divisor)) {
			continue
		}

		maxDivisor = divisor
	}

	return maxDivisor
}

func checkWithDivisor(str1, divisor string) bool {
	for i := 0; i < len(str1); i += len(divisor) {
		if i+len(divisor) > len(str1) {
			return false
		}

		if str1[i:i+len(divisor)] != divisor {
			return false
		}
	}
	return true
}
