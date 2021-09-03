package main

import (
	"fmt"
	"math"
)

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 0}

	fmt.Println(getDecimalValue(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	out := 0
	for head != nil {
		out = out<<1 | head.Val
		head = head.Next
	}

	return out
}

// straightforward solution, not optimal
func getDecimalValue1(head *ListNode) int {
	vals := make([]int, 0)
	node := head
	for node != nil {
		vals = append(vals, node.Val)
		node = node.Next
	}

	result := 0
	var pow float64 = 0
	for i := len(vals) - 1; i >= 0; i-- {
		result += int(math.Pow(2, pow)) * vals[i]
		pow++
	}

	return result
}
