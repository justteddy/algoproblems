package main

import "fmt"

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(isArmstrong(153))
	fmt.Println(isArmstrong(2))
}

func isArmstrong(n int) bool {
	sum, dig, num := 0, 0, n

	for n > 0 {
		dig++
		n /= 10
	}

	n = num

	for n != 0 {
		t := n % 10
		sum += power(t, dig)

		if sum > num {
			return false
		}
		n /= 10
	}

	return sum == num
}

func power(n, p int) int {
	if p == 0 {
		return 1
	}
	result := 1
	for i := 0; i < p; i++ {
		result *= n
	}
	return result
}
