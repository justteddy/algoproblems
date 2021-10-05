package main

import "fmt"

func main() {
	fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	curr := num
	for i := 0; curr > 0; i++ {
		if curr&1 == 1 {
			mask := ^(1 << i)
			num &= mask
		} else {
			num |= (1 << i)
		}
		curr >>= 1
	}
	return num
}
