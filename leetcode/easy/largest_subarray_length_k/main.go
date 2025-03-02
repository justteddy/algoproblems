package largest_subarray_length_k

// https://leetcode.com/problems/largest-subarray-length-k/
func largestSubarray(nums []int, k int) []int {
	idx := 0
	max := nums[0]
	for i := 1; i <= len(nums)-k; i++ {
		if nums[i] >= max {
			max = nums[i]
			idx = i
		}
	}

	return nums[idx : idx+k]
}
