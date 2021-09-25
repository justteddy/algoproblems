package main

import "fmt"

// https://leetcode.com/problems/middle-of-the-linked-list/
func main() {
	head := ListNode{Val: 0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(middleNode(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
