package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func main() {
	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head.Next
	prev := head
	for curr != nil {
		if curr.Val > prev.Val {
			prev.Next = curr
			prev = curr
		} else {
			prev.Next = nil
		}

		curr = curr.Next
	}

	return head
}
