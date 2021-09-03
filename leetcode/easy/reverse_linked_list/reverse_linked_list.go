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

// iterative solution
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
}

// recursive solution 1
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// recursive solution 2
func reverseList(head *ListNode) *ListNode {
	return reverse(head, nil)
}

func reverse(node, prev *ListNode) *ListNode {
	if node == nil {
		return prev
	}

	next := node.Next
	node.Next = prev
	return reverse(next, node)
}
