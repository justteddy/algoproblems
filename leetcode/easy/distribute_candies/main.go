package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	limit := len(candyType) / 2
	unique := make(map[int]struct{})
	for _, candy := range candyType {
		unique[candy] = struct{}{}
	}

	if len(unique) <= limit {
		return len(unique)
	}

	return limit
}
