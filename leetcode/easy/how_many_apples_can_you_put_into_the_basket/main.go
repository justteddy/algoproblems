package main

import "sort"

func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)

	cnt, sum := 0, 0
	for _, v := range weight {
		sum += v
		if sum <= 5000 {
			cnt++
			continue
		}
		break
	}

	return cnt
}
