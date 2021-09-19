package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/determine-color-of-a-chessboard-square/
func main() {
	fmt.Println(squareIsWhite("a1"))
}

// ASCII
// [1...8] - [49...56]
// [a...h] - [97...104]
// we can just sum their ASCII codes and get the same result
func squareIsWhite(coordinates string) bool {
	return (coordinates[0]+coordinates[1])&1 == 1
}

func squareIsWhite_(coordinates string) bool {
	dict := make(map[uint8]int)
	for i := 0; i < 8; i = i + 1 {
		dict[uint8(97+i)] = i + 1
	}

	x := dict[coordinates[0]]
	y, _ := strconv.Atoi(string(coordinates[1]))

	if (x+y)&1 == 0 {
		return false
	}
	return true
}
