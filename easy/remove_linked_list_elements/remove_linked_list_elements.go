package remove_linked_list_elements

/*
   Remove all elements from a linked list of integers that have value val.

   Example:
   Input:  1->2->6->3->4->5->6, val = 6
   Output: 1->2->3->4->5

   Input:  d->2->2->2->3->4->5->6, val = 2
   Output: 3->4->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	return dummy.Next
}
