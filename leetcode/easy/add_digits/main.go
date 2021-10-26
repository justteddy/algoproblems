package main

import "fmt"

// https://leetcode.com/problems/add-digits/
func main() {
	fmt.Println(addDigits(12334124))
}

func addDigits(num int) int {
	result := 0
	for num != 0 {
		result += num % 10
		num /= 10
	}
	if result >= 10 {
		return addDigits(result)
	}
	return result
}
