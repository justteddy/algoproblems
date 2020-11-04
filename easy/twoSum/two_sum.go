package two_sum

/*
   Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
   You may assume that each input would have exactly one solution, and you may not use the same element twice.
   You can return the answer in any order.

   Example:
   Input: nums = [2,7,11,15], target = 9
   Output: [0,1]
   Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, ok := dict[target-num]; ok {
			return []int{i, dict[target-num]}
		}
		dict[num] = i
	}
	return nil
}
