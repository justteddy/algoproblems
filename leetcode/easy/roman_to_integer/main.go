package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	hm := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := hm[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if hm[s[i]] < hm[s[i+1]] {
			result -= hm[s[i]]
		} else {
			result += hm[s[i]]
		}
	}

	return result
}

// fun and complex solution
func romanToInt_(s string) int {
	hm := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result, prev := 0, ""
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			if prev != "" {
				tmp := hm[prev[len(prev)-1]]
				for j := len(prev) - 2; j >= 0; j-- {
					tmp -= hm[prev[j]]
				}
				result += tmp - hm[s[i]]
				prev = ""
				continue
			}
			result += hm[s[i]]
			continue
		}

		if hm[s[i-1]] >= hm[s[i]] {
			if prev != "" {
				tmp := hm[prev[len(prev)-1]]
				for j := len(prev) - 2; j >= 0; j-- {
					tmp -= hm[prev[j]]
				}
				result += tmp - hm[s[i]]
				prev = ""
				continue
			}
			result += hm[s[i]]
			continue
		}
		prev += string(s[i])
	}

	return result
}
