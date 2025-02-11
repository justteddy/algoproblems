package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func longestCommonPrefix2(strs []string) string {
	prefix := strs[0] // flower
	for i, s := range strs {
		if i == 0 {
			continue
		}
		if len(s) == 0 || prefix == "" {
			return ""
		}

		newPrefix := prefix
		for j := 0; j < len(prefix); j++ {
			if j >= len(s) || s[j] != prefix[j] {
				newPrefix = s[:j]
				break
			}
			if s[j] == prefix[j] {
				newPrefix = s[:j+1]
			}
		}
		prefix = newPrefix
	}
	return prefix
}
