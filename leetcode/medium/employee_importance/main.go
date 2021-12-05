package main

import "fmt"

// https://leetcode.com/problems/employee-importance/
func main() {
	fmt.Println(getImportance([]*Employee{
		{
			Id:           1,
			Importance:   5,
			Subordinates: []int{2, 3},
		},
		{
			Id:           2,
			Importance:   3,
			Subordinates: nil,
		},
		{
			Id:           3,
			Importance:   3,
			Subordinates: nil,
		},
	}, 1))
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// recursive solution
func getImportance(employees []*Employee, id int) int {
	dict := make(map[int]*Employee)
	for _, employee := range employees {
		dict[employee.Id] = employee
	}

	result := 0
	var traverse func(int)
	traverse = func(ident int) {
		result += dict[ident].Importance
		for _, sub := range dict[ident].Subordinates {
			traverse(sub)
		}
	}

	traverse(id)

	return result
}

// iterative solution
func getImportance_(employees []*Employee, id int) int {
	dict := make(map[int]*Employee)
	for _, employee := range employees {
		dict[employee.Id] = employee
	}

	result := 0
	stack := make([]int, 0)
	stack = append(stack, id)

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result += dict[curr].Importance
		for _, sub := range dict[curr].Subordinates {
			stack = append(stack, sub)
		}
	}

	return result
}
