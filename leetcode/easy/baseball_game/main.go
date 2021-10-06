package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

func calPoints(ops []string) int {
	stack := new(Stack)
	for _, op := range ops {
		switch op {
		case "C":
			stack.Pop()
		case "D":
			prev := stack.Pop()
			stack.Push(prev)
			stack.Push(prev * 2)
		case "+":
			prev1 := stack.Pop()
			prev2 := stack.Pop()
			stack.Push(prev2)
			stack.Push(prev1)
			stack.Push(prev1 + prev2)
		default:
			val, _ := strconv.Atoi(op)
			stack.Push(val)
		}
	}

	result := 0
	for !stack.IsEmpty() {
		result += stack.Pop()
	}

	return result
}

type Stack []int

func (s *Stack) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack) Len() int      { return len(*s) }
func (s *Stack) Push(el int)   { *s = append(*s, el) }
func (s *Stack) Pop() int {
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}
