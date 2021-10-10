package height_checker

import "sort"

/*
   Students are asked to stand in non-decreasing order of heights for an annual photo.
   Return the minimum number of students that must move in order for all students to be standing in non-decreasing order of height.
   Notice that when a group of students is selected they can reorder in any possible way between themselves and the non selected students remain on their seats.

   Example 1:
   Input: heights = [1,1,4,2,1,3]
   Output: 3

   Explanation:
   Current array : [1,1,4,2,1,3]
   Target array  : [1,1,1,2,3,4]
   On index 2 (0-based) we have 4 vs 1 so we have to move this student.
   On index 4 (0-based) we have 1 vs 3 so we have to move this student.
   On index 5 (0-based) we have 3 vs 4 so we have to move this student.

   Example 2:
   Input: heights = [5,1,2,3,4]
   Output: 5

   Example 3:
   Input: heights = [1,2,3,4,5]
   Output: 0
*/

func heightChecker_(heights []int) int {
	s := make([]int, len(heights))
	copy(s, heights)
	sort.Ints(s)

	cnt := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}

func heightChecker(heights []int) int {
	count := make([]int, 101)
	for i := 0; i < len(heights); i++ {
		count[heights[i]]++
	}
	ans := 0
	for i, j := 1, 0; i < len(count); i++ {
		for k := count[i]; k > 0; k-- {
			if i != heights[j] {
				ans++
			}
			j++
		}
	}
	return ans
}
