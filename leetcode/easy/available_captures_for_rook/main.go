package main

import "fmt"

// https://leetcode.com/problems/available-captures-for-rook/
func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}

func numRookCaptures(board [][]byte) int {
	rookRow, rookCol := findRook(board)
	up, down, left, right := rookRow-1, rookRow+1, rookCol-1, rookCol+1
	result := 0

	// up
UP:
	for up >= 0 {
		switch board[up][rookCol] {
		case '.':
			up--
			continue
		case 'B':
			break UP
		case 'p':
			result++
			break UP
		}
	}

	// down
DOWN:
	for down < len(board) {
		switch board[down][rookCol] {
		case '.':
			down++
			continue
		case 'B':
			break DOWN
		case 'p':
			result++
			break DOWN
		}
	}

	// left
LEFT:
	for left >= 0 {
		switch board[rookRow][left] {
		case '.':
			left--
			continue
		case 'B':
			break LEFT
		case 'p':
			result++
			break LEFT
		}
	}

	// right
RIGHT:
	for right < len(board[0]) {
		switch board[rookRow][right] {
		case '.':
			right++
			continue
		case 'B':
			break RIGHT
		case 'p':
			result++
			break RIGHT
		}
	}

	return result
}

func findRook(board [][]byte) (int, int) {
	for r, rows := range board {
		for c, col := range rows {
			if col == 'R' {
				return r, c
			}
		}
	}
	return 0, 0
}
