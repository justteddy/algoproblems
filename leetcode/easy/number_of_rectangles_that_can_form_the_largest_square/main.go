package main

import "fmt"

func main() {
	fmt.Println(countGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {5, 14}}))
}

// O(N) solution with 0(1) space
func countGoodRectangles(rectangles [][]int) int {
	maxSize, count := 0, 0
	for _, sizes := range rectangles {
		squareSide := min(sizes[0], sizes[1])
		if squareSide == maxSize {
			count++
			continue
		}
		if squareSide > maxSize {
			maxSize = squareSide
			count = 1
			continue
		}
	}

	return count
}

// O(N) solution with 0(N) space - hashmap
func countGoodRectangles_(rectangles [][]int) int {
	hm := make(map[int]int)
	for _, sizes := range rectangles {
		curr := min(sizes[0], sizes[1])
		if _, ok := hm[curr]; !ok {
			hm[curr] = 1
			continue
		}
		hm[curr]++
	}

	max, ans := 0, 0
	for size, count := range hm {
		if size > max {
			max = size
			ans = count
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
