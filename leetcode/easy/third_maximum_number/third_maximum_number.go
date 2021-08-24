package third_maximum_number

import "sort"

/*
   Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

   Example 1:
   Input: [3, 2, 1]
   Output: 1
   Explanation: The third maximum is 1.

   Example 2:
   Input: [1, 2]
   Output: 2
   Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

   Example 3:
   Input: [2, 2, 3, 1]
   Output: 1
   Explanation: Note that the third maximum here means the third maximum distinct number.
   Both numbers with value 2 are both considered as second maximum.

*/

func thirdMax(nums []int) int {
	unique := make([]int, 0)
	dict := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := dict[num]; !ok {
			dict[num] = struct{}{}
			unique = append(unique, num)
		}
	}

	sort.Ints(unique)

	if len(unique) < 3 {
		return unique[len(unique)-1]
	}

	return unique[len(unique)-3]
}
