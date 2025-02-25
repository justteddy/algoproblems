package main

import (
	"fmt"
	"math"
)

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println("min - ", ms.GetMin())
	ms.Pop()
	fmt.Println("top - ", ms.Top())
	fmt.Println("min - ", ms.GetMin())
}

type MinStack struct {
	els [][2]int
	min int
}

func Constructor() MinStack {
	return MinStack{
		els: make([][2]int, 0, 1000),
		min: math.MaxInt32,
	}
}

func (s *MinStack) Push(val int) {
	if val < s.min {
		s.min = val
	}

	s.els = append(s.els, [2]int{val, s.min})
}

func (s *MinStack) Pop() {
	if len(s.els) == 0 {
		return
	}

	s.els = s.els[:len(s.els)-1]
	if len(s.els) == 0 {
		s.min = math.MaxInt32
	} else {
		s.min = s.els[len(s.els)-1][1]
	}
}

func (s *MinStack) Top() int {
	return s.els[len(s.els)-1][0]
}

func (s *MinStack) GetMin() int {
	return s.min
}
