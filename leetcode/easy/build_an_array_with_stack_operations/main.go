package main

import "fmt"

// https://leetcode.com/problems/build-an-array-with-stack-operations/
func main() {
	fmt.Println(buildArray([]int{1, 3}, 4))
}

func buildArray(target []int, n int) []string {
	result := make([]string, 0)
	for i, j := 1, 0; i <= n; i++ {
		if j == len(target) {
			break
		}
		if target[j] == i {
			result = append(result, "Push")
			j++
		} else {
			result = append(result, []string{"Push", "Pop"}...)
		}
	}

	return result
}
