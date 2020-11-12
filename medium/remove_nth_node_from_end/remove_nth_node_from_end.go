package remove_nth_node_from_end

/*
   Given the head of a linked list, remove the nth node from the end of the list and return its head.
   Follow up: Could you do this in one pass?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	fast, slow := head, head
	for i := 0; fast != nil; i++ {
		slow = slow.Next
		fast = fast.Next
		if i == n {
			slow = head
		}
	}
	slow.Next = slow.Next.Next
	return head.Next
}
