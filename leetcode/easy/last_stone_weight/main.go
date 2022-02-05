package main

import "fmt"

// https://leetcode.com/problems/last-stone-weight/
func main() {
	fmt.Println(lastStoneWeight([]int{3, 1}))
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	for {
		maxIdx, prevMaxIdx := findTwoHeaviestIdx(stones)
		if stones[maxIdx] == 0 {
			return stones[prevMaxIdx]
		}

		if stones[prevMaxIdx] == 0 {
			return stones[maxIdx]
		}

		if stones[maxIdx] == stones[prevMaxIdx] {
			stones[maxIdx] = 0
			stones[prevMaxIdx] = 0
		} else {
			stones[maxIdx] = stones[maxIdx] - stones[prevMaxIdx]
			stones[prevMaxIdx] = 0
		}
	}
}

func findTwoHeaviestIdx(stones []int) (int, int) {
	max, maxIdx := -1, 0
	for idx, weight := range stones {
		if weight > max {
			max = weight
			maxIdx = idx
		}
	}

	prevMax, prevMaxIdx := -1, 0
	for idx, weight := range stones {
		if idx == maxIdx {
			continue
		}

		if weight > prevMax {
			prevMax = weight
			prevMaxIdx = idx
		}
	}

	return maxIdx, prevMaxIdx
}
