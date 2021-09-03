package main

import "fmt"

/*
	Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.
	A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:
	0 <= i < j < k < arr.length
	|arr[i] - arr[j]| <= a
	|arr[j] - arr[k]| <= b
	|arr[i] - arr[k]| <= c
	Where |x| denotes the absolute value of x.
	Return the number of good triplets.

	Example 1:
	Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
	Output: 4
	Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].

	Example 2:
	Input: arr = [1,1,2,2,3], a = 0, b = 0, c = 1
	Output: 0
	Explanation: No triplet satisfies all conditions.

	Constraints:
	3 <= arr.length <= 100
	0 <= arr[i] <= 1000
	0 <= a, b, c <= 1000
*/

func main() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if match(arr[i], arr[j], a) && match(arr[j], arr[k], b) && match(arr[i], arr[k], c) {
					result++
				}
			}
		}
	}

	return result
}

func match(n1, n2, value int) bool {
	return abs(n1-n2) <= value
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
