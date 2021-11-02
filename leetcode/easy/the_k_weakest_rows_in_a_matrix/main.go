package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3))
}

func kWeakestRows(mat [][]int, k int) []int {
	type rowMeta struct {
		idx   int
		power int
	}

	counts := make([]rowMeta, 0, len(mat))
	for i, row := range mat {
		power := 0
		for _, unit := range row {
			if unit != 1 {
				break
			}
			power++
		}
		counts = append(counts, rowMeta{idx: i, power: power})
	}

	sort.Slice(counts, func(i, j int) bool {
		if counts[i].power == counts[j].power {
			return counts[i].idx < counts[j].idx
		}

		return counts[i].power < counts[j].power
	})

	counts = counts[:k]
	result := make([]int, 0, len(counts))
	for _, count := range counts {
		result = append(result, count.idx)
	}

	return result
}
