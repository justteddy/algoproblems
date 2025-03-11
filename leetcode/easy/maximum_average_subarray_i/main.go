package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

// sliding window elegant
func findMaxAverage(nums []int, k int) float64 {
	mx := float64(math.MinInt64)
	slidingSum := 0

	for i := 0; i < len(nums); i++ {
		slidingSum += nums[i]
		if i >= k-1 { // i>= 3
			if i > k-1 {
				slidingSum -= nums[i-k]
			}
			s := float64(slidingSum) / float64(k)
			if s > mx {
				mx = s
			}
		}
	}

	return mx
}

// brute force solution
func findMaxAverage2(nums []int, k int) float64 {
	avgFunc := func(from int) float64 {
		sum := 0
		for _, v := range nums[from : from+k] {
			sum += v
		}

		return float64(sum) / float64(k)
	}

	mx := avgFunc(0)
	for i := 0; i <= len(nums)-k; i++ {
		s := avgFunc(i)
		if s > mx {
			mx = s
		}
	}

	return mx
}
