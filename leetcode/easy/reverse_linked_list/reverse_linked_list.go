package reverse_linked_list

/*
   Reverse a singly linked list.

   Example:
   Input: 1->2->3->4->5->NULL
   Output: 5->4->3->2->1->NULL

   Follow up:
   A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	crt := head
	var prev *ListNode

	for crt != nil {
		nextTmp := crt.Next
		crt.Next = prev
		prev = crt
		crt = nextTmp
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
