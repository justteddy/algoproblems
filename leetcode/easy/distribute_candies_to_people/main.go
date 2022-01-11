package main

import "fmt"

// https://leetcode.com/problems/distribute-candies-to-people/
func main() {
	fmt.Println(distributeCandies(7, 4))
}

func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)
	i, portion := 0, 1
	for candies > 0 {
		if candies < portion {
			result[i] += candies
			break
		}

		result[i] += portion
		candies -= portion

		portion++
		i++
		if i == len(result) {
			i = 0
		}
	}

	return result
}
