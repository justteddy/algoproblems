package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("   hello     world   "))
}

func reverseWords2(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWords(s string) string {
	b := []byte(s)
	b = trimSpaces(b)
	b = cleanSpaces(b)

	n := len(b)
	var res []byte
	end := n

	for i := n - 1; i >= 0; i-- {
		if b[i] == ' ' {
			res = append(res, b[i+1:end]...)
			res = append(res, ' ')
			end = i
		}
	}
	res = append(res, b[:end]...)
	return string(res)
}

func trimSpaces(b []byte) []byte {
	start, end := 0, len(b)-1
	for start <= end && b[start] == ' ' {
		start++
	}
	for end >= start && b[end] == ' ' {
		end--
	}
	return b[start : end+1]
}

func cleanSpaces(s []byte) []byte {
	result := make([]byte, 0, len(s))
	seqSpaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			seqSpaces++
			continue
		}
		if seqSpaces > 0 {
			result = append(result, 32)
			seqSpaces = 0
		}

		result = append(result, s[i])
	}

	return result
}
