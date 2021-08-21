package main

import "fmt"

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) // leetcode
	fmt.Println(restoreString("aiohn", []int{3, 1, 4, 2, 0}))             // nihao
}

func restoreString(s string, indices []int) string {
	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ans[indices[i]] = s[i]
	}
	return string(ans)
}
