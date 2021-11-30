package main

import "fmt"

// https://leetcode.com/problems/nim-game/
func main() {
	fmt.Println(canWinNim(16))
}

func canWinNim(n int) bool {
	return n%4 != 0
}
