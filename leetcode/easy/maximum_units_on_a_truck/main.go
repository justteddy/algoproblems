package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-units-on-a-truck/
func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] < boxTypes[j][1]
	})

	units := 0
	for i := len(boxTypes) - 1; i >= 0; i-- {
		if boxTypes[i][0] > truckSize {
			units += truckSize * boxTypes[i][1]
			break
		} else {
			truckSize -= boxTypes[i][0]
			units += boxTypes[i][0] * boxTypes[i][1]
			continue
		}
	}

	return units
}
