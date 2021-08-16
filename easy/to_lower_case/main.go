package main

import "fmt"

/*
	Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.

	Example 1:
	Input: s = "Hello"
	Output: "hello"

	Example 2:
	Input: s = "here"
	Output: "here"

	Example 3:
	Input: s = "LOVELY"
	Output: "lovely"

	Constraints:
	1 <= s.length <= 100
	s consists of printable ASCII characters.
*/
func main() {
	fmt.Println(toLowerCase("Hello"))
}

func toLowerCase(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > 96 {
			result[i] = s[i]
			continue
		}
		result[i] = s[i] + 32
	}

	return string(result)
}
