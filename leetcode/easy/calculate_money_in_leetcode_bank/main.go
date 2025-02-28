package main

import "fmt"

func main() {
	fmt.Println(totalMoney(20))
}

// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
func totalMoney(n int) int {
	res, startVal := 0, 0
	curr := startVal
	for i := 1; i <= n; i++ {
		if i%7 == 1 {
			startVal++
			curr = startVal
		}
		res += curr
		curr++
	}
	return res
}
