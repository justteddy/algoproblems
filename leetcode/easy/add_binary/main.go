package main

import (
	"fmt"
)

// https://leetcode.com/problems/add-binary/
func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return calc(a, b)
	}

	return calc(b, a)
}

func calc(first, second string) string {
	result := make([]byte, 0, len(first))
	var debt bool
	fIdx, sIdx := len(first)-1, len(second)-1
	for fIdx >= 0 {
		f, s := 0, 0
		if first[fIdx] == 49 {
			f = 1
		}
		if sIdx >= 0 && second[sIdx] == 49 {
			s = 1
		}
		switch f + s {
		case 2:
			if debt {
				result = append(result, '1')
			} else {
				result = append(result, '0')
				debt = true
			}
		case 1:
			if debt {
				result = append(result, '0')
			} else {
				result = append(result, '1')
			}
		default:
			if debt {
				result = append(result, '1')
				debt = false
			} else {
				result = append(result, '0')
			}
		}
		fIdx--
		sIdx--
	}

	if debt {
		result = append(result, '1')
	}

	// reverse slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
