package main

import "fmt"

// https://leetcode.com/problems/matrix-cells-in-distance-order/
func main() {
	fmt.Println(allCellsDistOrder(4, 4, 2, 3))
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	result := make([][]int, 0)

	seen := make(map[string]struct{})
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{rCenter, cCenter})
	for len(queue) != 0 {
		cell := queue[0]
		queue = queue[1:]

		if cell[0] < 0 || cell[0] >= rows || cell[1] < 0 || cell[1] >= cols {
			continue
		}

		k := key(cell[0], cell[1])

		if _, ok := seen[k]; ok {
			continue
		}

		result = append(result, []int{cell[0], cell[1]})
		seen[k] = struct{}{}

		for _, dir := range dirs {
			queue = append(queue, [2]int{cell[0] + dir[0], cell[1] + dir[1]})
		}
	}

	return result
}

func key(r, c int) string {
	return fmt.Sprintf("%d.%d", r, c)
}
