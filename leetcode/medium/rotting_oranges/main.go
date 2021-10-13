package main

import "fmt"

// https://leetcode.com/problems/rotting-oranges/
func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}

func orangesRotting(grid [][]int) int {
	queue := new(Queue)
	fresh := 0

	for i, rows := range grid {
		for j, col := range rows {
			if col == 2 {
				queue.Enqueue([2]int{i, j})
			} else if col == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for !queue.IsEmpty() && fresh > 0 {
		minutes++

		levelLen := queue.Len()
		for levelLen > 0 {
			rotten := queue.Dequeue()
			y, x := rotten[0], rotten[1]

			// up
			if adjacentFresh(y-1, x, grid) {
				grid[y-1][x] = 2
				queue.Enqueue([2]int{y - 1, x})
				fresh--
			}
			// down
			if adjacentFresh(y+1, x, grid) {
				grid[y+1][x] = 2
				queue.Enqueue([2]int{y + 1, x})
				fresh--
			}
			// left
			if adjacentFresh(y, x-1, grid) {
				grid[y][x-1] = 2
				queue.Enqueue([2]int{y, x - 1})
				fresh--
			}
			// right
			if adjacentFresh(y, x+1, grid) {
				grid[y][x+1] = 2
				queue.Enqueue([2]int{y, x + 1})
				fresh--
			}

			levelLen--
		}
	}

	if fresh == 0 {
		return minutes
	}

	return -1
}

type Queue [][2]int

func (q *Queue) IsEmpty() bool     { return len(*q) == 0 }
func (q *Queue) Len() int          { return len(*q) }
func (q *Queue) Enqueue(el [2]int) { *q = append(*q, el) }
func (q *Queue) Dequeue() [2]int {
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

func adjacentFresh(y, x int, grid [][]int) bool {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return false
	}

	if grid[y][x] == 0 || grid[y][x] == 2 {
		return false
	}

	return true
}
