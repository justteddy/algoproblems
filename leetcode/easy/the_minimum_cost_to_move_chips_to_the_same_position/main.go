package main

import "fmt"

// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 2, 2, 2, 3, 4, 5, 55, 6}))
}

func minCostToMoveChips(position []int) int {
	oddCnt, evenCnt := 0, 0
	for _, val := range position {
		if val&1 == 1 {
			oddCnt++
		} else {
			evenCnt++
		}
	}

	if oddCnt < evenCnt {
		return oddCnt
	}
	return evenCnt
}
