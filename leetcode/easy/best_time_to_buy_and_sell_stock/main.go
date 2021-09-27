package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 9, 5, 2, 0, 6}))
}

// simple
func maxProfit(prices []int) int {
	min, maxProfit := 10000, 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > maxProfit {
			maxProfit = v - min
		}
	}
	return maxProfit
}

// cache results - dynamic programming
func maxProfit_(prices []int) int {
	max := 0
	cache := make(map[int]int) // TODO: allocate once
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		cache[i-1] = max
	}

	result := 0
	for i, v := range prices {
		if cache[i]-v > result {
			result = cache[i] - v
		}
	}

	return result
}
