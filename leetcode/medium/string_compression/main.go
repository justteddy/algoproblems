package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/string-compression
func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	n := len(chars)
	write := 0
	for i := 0; i < n; {
		current := chars[i]
		count := 0
		for i < n && chars[i] == current {
			i++
			count++
		}
		chars[write] = current
		write++
		if count > 1 {
			countStr := strconv.Itoa(count)
			for j := 0; j < len(countStr); j++ {
				chars[write] = countStr[j]
				write++
			}
		}
	}
	return write
}
