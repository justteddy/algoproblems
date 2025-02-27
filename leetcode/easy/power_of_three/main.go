package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
}

func isPowerOfThree(n int) bool {
	if n < 3 || n%3 != 0 {
		return false
	}

	for i := 3; i < n; i *= 3 {
		if i == n {
			return true
		}
	}

	return false
}

func isPowerOfThree2(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}
