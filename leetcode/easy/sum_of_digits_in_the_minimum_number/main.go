package main

import "fmt"

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8}))
}

func sumOfDigits(nums []int) int {
	min := 100
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	sum := sumOfDigitsInInt(min)

	if sum&1 == 1 {
		return 0
	}
	return 1
}

func sumOfDigitsInInt(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
