package main

// https://leetcode.com/problems/valid-sudoku/
func main() {

}

func isValidSudoku(board [][]byte) bool {
	// check rows and cols
	for i := 0; i < 9; i++ {
		row := make(map[byte]struct{})
		col := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			// check row
			if board[i][j] != '.' {
				if _, ok := row[board[i][j]]; ok {
					return false
				}
				row[board[i][j]] = struct{}{}
			}

			// check col
			if board[j][i] != '.' {
				if _, ok := col[board[j][i]]; ok {
					return false
				}
				col[board[j][i]] = struct{}{}
			}
		}
	}

	// check boxes
	for _, rowStart := range []int{0, 3, 6} {
		for _, colStart := range []int{0, 3, 6} {
			if !checkBox(board, rowStart, colStart) {
				return false
			}
		}
	}

	return true
}

func checkBox(board [][]byte, rowStart, colStart int) bool {
	m := make(map[byte]struct{})
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := m[board[i][j]]; ok {
				return false
			}
			m[board[i][j]] = struct{}{}
		}
	}
	return true
}
