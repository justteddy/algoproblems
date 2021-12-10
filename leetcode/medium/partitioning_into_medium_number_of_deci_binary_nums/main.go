package main

import (
	"fmt"
)

// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
func main() {
	fmt.Println(minPartitions("4957"))
}

func minPartitions(n string) int {
	max := 0
	for i, _ := range n {
		if int(n[i]-'0') > max {
			max = int(n[i] - '0')
		}
	}
	return max
}
