package max_consecutive_ones

/*
	Given a binary array, find the maximum number of consecutive 1s in this array.

	Example:
	Input: [1,1,0,1,1,1]
	Output: 3
	Explanation: The first two digits or the last three digits are consecutive 1s.
		The maximum number of consecutive 1s is 3.

	Note:
		The input array will only contain 0 and 1.
		The length of input array is a positive integer and will not exceed 10,000

*/

func findMaxConsecutiveOnes(nums []int) int {
	max, current := 0, 0
	for _, num := range nums {
		if current+num > max {
			max = current + num
		}
		current = (current + num) * num
	}
	return max
}
