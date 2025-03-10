package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 1, 0, 0, 1, 0, 0}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	prev, next := 0, 0
	for i, v := range flowerbed {
		if i+1 < len(flowerbed) {
			next = flowerbed[i+1]
		}

		if i-1 >= 0 {
			prev = flowerbed[i-1]
		}

		if next == 0 && prev == 0 && v != 1 {
			flowerbed[i] = 1
			n--
		}

		if n == 0 {
			return true
		}
	}

	return false
}
