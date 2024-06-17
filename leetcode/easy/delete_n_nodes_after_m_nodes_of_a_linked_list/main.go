package main

import "fmt"

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println(deleteNodes(&head, 2, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := ""
	for l != nil {
		res += fmt.Sprintf("%v -> ", l.Val)
		l = l.Next
	}
	return res
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	keep, del := m, n
	node := head
	var prev *ListNode
	for node != nil {
		if keep > 0 {
			prev = node
			node = node.Next
			keep--
			continue
		}
		if del > 0 {
			node = node.Next
			del--
			continue
		}
		keep, del = m, n
		prev.Next = node
	}

	if del < n {
		prev.Next = nil
	}

	return head
}
