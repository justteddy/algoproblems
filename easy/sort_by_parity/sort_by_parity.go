package sort_by_parity

/*
   Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
   You may return any answer array that satisfies this condition.

   Example 1:
   Input: [3,1,2,4]
   Output: [2,4,3,1]
   The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

*/

func sortArrayByParity(nums []int) []int {
	wrtPtr := 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i]%2 == 0 {
			nums[wrtPtr], nums[i] = nums[i], nums[wrtPtr]
			wrtPtr++
		}
	}
	return nums
}
