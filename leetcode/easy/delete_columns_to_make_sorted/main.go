package main

import "fmt"

// https://leetcode.com/problems/delete-columns-to-make-sorted/
func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}

func minDeletionSize(strs []string) int {
	ans := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				ans++
				break
			}
		}
	}

	return ans
}
