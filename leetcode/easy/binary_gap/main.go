package main

import "fmt"

// https://leetcode.com/problems/binary-gap/
func main() {
	fmt.Println(binaryGap(22))
}

func binaryGap(n int) int {
	longest := 0
	for n != 0 {
		bit := n & 1
		if bit == 0 {
			n >>= 1
			continue
		}
		distance := distanceToNearestOneBit(n >> 1)
		if distance == 0 {
			break
		}
		if distance > longest {
			longest = distance
		}

		n >>= 1
	}

	return longest
}

func distanceToNearestOneBit(n int) int {
	distance := 1
	for n != 0 {
		if n&1 == 1 {
			return distance
		}
		distance++
		n >>= 1
	}

	return 0
}
