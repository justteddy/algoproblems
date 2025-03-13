package main

import "fmt"

func main() {
	fmt.Println(deleteMiddle(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
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

// fast and slow pointers
func deleteMiddle(head *ListNode) *ListNode {
	fast, slow, prevSlow := head, head, head
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow == head {
		return nil
	}

	prevSlow.Next = prevSlow.Next.Next
	return head
}

// 2 iterations
func deleteMiddle2(head *ListNode) *ListNode {
	cnt := 0
	n := head
	for n != nil {
		n = n.Next
		cnt++
	}

	if cnt == 1 {
		return nil
	}

	middle := cnt / 2
	n = head
	cnt = 0
	for {
		if cnt == middle-1 {
			n.Next = n.Next.Next
			break
		}

		n = n.Next
		cnt++
	}

	return head
}
