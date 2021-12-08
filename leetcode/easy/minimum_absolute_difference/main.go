package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-absolute-difference/
func main() {
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	m := make(map[int][][]int)

	smallest := 1_000_000
	i, j := 0, 1
	for j < len(arr) {
		diff := absDiff(arr[i], arr[j])
		if diff <= smallest {
			smallest = diff
			if _, ok := m[smallest]; !ok {
				m[smallest] = make([][]int, 0)
			}
			if arr[i] < arr[j] {
				m[smallest] = append(m[smallest], []int{arr[i], arr[j]})
			} else {
				m[smallest] = append(m[smallest], []int{arr[j], arr[i]})
			}
		}

		i, j = i+1, j+1
	}

	diff := 1
	for {
		if _, ok := m[diff]; !ok {
			diff++
			continue
		}
		sort.Slice(m[diff], func(i, j int) bool {
			return m[diff][i][0]+m[diff][j][1] < m[diff][j][0]+m[diff][j][1]
		})
		break
	}

	return m[diff]
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}

	return diff
}
