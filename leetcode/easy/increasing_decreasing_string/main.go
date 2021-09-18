package main

import (
	"fmt"
)

// https://leetcode.com/problems/increasing-decreasing-string/
func main() {
	fmt.Println(sortString("leetcode"))
}

func sortString(s string) string {
	d := dict(s)
	result := make([]byte, 0, len(s))
	for len(d) > 0 {
		var prev uint8
		for i := uint8(97); i <= 122; i++ {
			_, ok := d[i]
			if prev != i && ok {
				prev = i
				result = append(result, i)
				d[i]--
				if d[i] == 0 {
					delete(d, i)
				}
			}
		}

		prev = 0
		for i := uint8(122); i >= 97; i-- {
			_, ok := d[i]
			if prev != i && ok {
				prev = i
				result = append(result, i)
				d[i]--
				if d[i] == 0 {
					delete(d, i)
				}
			}
		}
	}

	return string(result)
}

func dict(s string) map[uint8]int {
	hm := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := hm[s[i]]; !ok {
			hm[s[i]] = 1
			continue
		}
		hm[s[i]]++
	}
	return hm
}
