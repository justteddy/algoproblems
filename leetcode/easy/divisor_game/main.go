package main

import "fmt"

// https://leetcode.com/problems/divisor-game/
func main() {
	fmt.Println(divisorGame(101))
}

func divisorGame(n int) bool {
	if n&1 == 1 {
		return false
	}
	return true
}
