package main

import "fmt"

// https://leetcode.com/problems/reverse-bits/
func main() {
	fmt.Printf("%b\n", reverseBits(43261596))
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
