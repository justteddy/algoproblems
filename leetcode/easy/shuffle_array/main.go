package main

import "fmt"

/*
	Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
	Return the array in the form [x1,y1,x2,y2,...,xn,yn].

	Example 1:
	Input: nums = [2,5,1,3,4,7], n = 3
	Output: [2,3,5,4,1,7]
	Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7].
	Example 2:

	Input: nums = [1,2,3,4,4,3,2,1], n = 4
	Output: [1,4,2,3,3,2,4,1]
	Example 3:

	Input: nums = [1,1,2,2], n = 2
	Output: [1,2,1,2]

	Constraints:
	1 <= n <= 500
	nums.length == 2n
	1 <= nums[i] <= 10^3
*/
func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}

func shuffle(nums []int, n int) []int {
	for i, j := len(nums)-1, n-1; i >= n; i, j = i-1, j-1 {
		nums[i] <<= 10
		nums[i] |= nums[j]
	}

	for i, j := 0, n; j < len(nums); i, j = i+2, j+1 {
		num1 := nums[j] & (1<<10 - 1)
		num2 := nums[j] >> 10
		nums[i] = num1
		nums[i+1] = num2
	}

	return nums
}

func shuffleSimple(nums []int, n int) []int {
	result := make([]int, 0, len(nums))
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		result = append(result, []int{nums[i], nums[j]}...)
	}
	return result
}
