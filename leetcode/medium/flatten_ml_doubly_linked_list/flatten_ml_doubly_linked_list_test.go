package flatten_ml_doubly_linked_list

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	/*
		   1---2---5---NULL
			   |
			   3---4---NULL
	*/
	t.Run("case 1", func(t *testing.T) {
		branch := Node{Val: 3}
		branch.Next = &Node{Val: 4, Prev: &branch}

		head := Node{Val: 1}
		head.Next = &Node{Val: 2, Prev: &head, Child: &branch}
		head.Next.Next = &Node{Val: 5, Prev: head.Next}

		traverse(flatten(&head)) // 1,2,3,4,5
	})

	/*
	   1---2---3---4---5---6--NULL
	           |
	           7---8---9---10--NULL
	               |
	               11--12--NULL
	*/
	t.Run("case 2", func(t *testing.T) {
		branch2 := Node{Val: 11}
		branch2.Next = &Node{Val: 12, Prev: &branch2}

		branch := Node{Val: 7}
		branch.Next = &Node{Val: 8, Prev: &branch, Child: &branch2}
		branch.Next.Next = &Node{Val: 9, Prev: branch.Next}
		branch.Next.Next.Next = &Node{Val: 10, Prev: branch.Next.Next}

		head := Node{Val: 1}
		head.Next = &Node{Val: 2, Prev: &head}
		head.Next.Next = &Node{Val: 3, Prev: head.Next, Child: &branch}
		head.Next.Next.Next = &Node{Val: 4, Prev: head.Next.Next}
		head.Next.Next.Next.Next = &Node{Val: 5, Prev: head.Next.Next.Next}
		head.Next.Next.Next.Next.Next = &Node{Val: 6, Prev: head.Next.Next.Next.Next}

		traverse(flatten(&head)) // 1,2,3,7,8,11,12,9,10,4,5,6
	})
}

func traverse(head *Node) {
	if head == nil {
		fmt.Println(nil)
		return
	}

	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Child != nil {
			traverse(cur.Child)
		}
		cur = cur.Next
	}
}
