package sorted_squares

/*
	Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.
	Example 1:
	Input: [-4,-1,0,3,10]
	Output: [0,1,9,16,100]

	Example 2:
	Input: [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
*/

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1

	for z := len(nums) - 1; z >= 0; z-- {
		sqi, sqj := nums[i]*nums[i], nums[j]*nums[j]

		if sqi > sqj {
			res[z] = sqi
			i++
			continue
		}
		res[z] = sqj
		j--
	}

	return res
}
