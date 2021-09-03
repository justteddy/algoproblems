package main

import "fmt"

/*
	Given an array of positive integers arr, calculate the sum of all possible odd-length subarrays.
	A subarray is a contiguous subsequence of the array.
	Return the sum of all odd-length subarrays of arr.

	Example 1:
	Input: arr = [1,4,2,5,3]
	Output: 58
	Explanation: The odd-length subarrays of arr and their sums are:
	[1] = 1
	[4] = 4
	[2] = 2
	[5] = 5
	[3] = 3
	[1,4,2] = 7
	[4,2,5] = 11
	[2,5,3] = 10
	[1,4,2,5,3] = 15
	If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58

	Example 2:
	Input: arr = [1,2]
	Output: 3
	Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.

	Example 3:
	Input: arr = [10,11,12]
	Output: 66

	Constraints:
	1 <= arr.length <= 100
	1 <= arr[i] <= 1000
*/
func main() {
	fmt.Println(sumOddLengthSubarrays([]int{0, 1, 2, 3, 4, 5}))
}

func sumOddLengthSubarrays(arr []int) int {
	result, subArrSize := 0, 1
	for subArrSize <= len(arr) {
		for i := 0; i+subArrSize <= len(arr); i++ {
			result += sum(arr[i : i+subArrSize])
		}
		subArrSize += 2
	}
	return result
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}
