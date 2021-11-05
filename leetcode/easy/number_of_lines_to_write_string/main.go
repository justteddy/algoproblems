package main

import "fmt"

// https://leetcode.com/problems/number-of-lines-to-write-string/
func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
}

func numberOfLines(widths []int, s string) []int {
	lines, currLineSize := 1, 0
	for i := 0; i < len(s); i++ {
		size := widths[int(s[i]-'a')]
		if currLineSize+size > 100 {
			lines++
			currLineSize = 0
		}
		currLineSize += size
	}

	return []int{lines, currLineSize}
}
