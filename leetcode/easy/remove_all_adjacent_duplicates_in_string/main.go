package main

import "fmt"

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(s string) string {
	stack := new(Stack)

	for i := 0; i < len(s); i++ {
		if !stack.IsEmpty() && stack.Peek() == s[i] {
			stack.Pop()
			continue
		}
		stack.Push(s[i])
	}

	return string(*stack)
}

type Stack []byte

func (s *Stack) IsEmpty() bool { return len(*s) == 0 }
func (s *Stack) Len() int      { return len(*s) }
func (s *Stack) Peek() byte    { return (*s)[len(*s)-1] }
func (s *Stack) Push(el byte)  { *s = append(*s, el) }
func (s *Stack) Pop() byte {
	element := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return element
}
