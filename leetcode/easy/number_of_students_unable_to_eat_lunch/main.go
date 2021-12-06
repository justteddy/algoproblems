package main

import "fmt"

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 1, 0}))
}

func countStudents(students []int, sandwiches []int) int {
	round := 0
	for round < len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			round = 0
			continue
		}

		student := students[0]
		students = students[1:]
		students = append(students, student)
		round++
	}

	return len(students)
}
