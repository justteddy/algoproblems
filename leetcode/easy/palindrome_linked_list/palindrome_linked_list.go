package palindrome_linked_list

/*
   Given a singly linked list, determine if it is a palindrome.

   Example 1:
   Input: 1->2
   Output: false

   Example 2:
   Input: 1->2->2->1
   Output: true

   Follow up:
   Could you do it in O(n) time and O(1) space?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	prev := &ListNode{}
	slowPtr, fastPtr := head, head
	for fastPtr != nil && fastPtr.Next != nil {
		prev = slowPtr
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	prev.Next = nil

	// if fast pointer not nil - it's odd count of elements
	// we need skip middle
	if fastPtr != nil {
		slowPtr = slowPtr.Next
	}

	head = reverse(head)
	for slowPtr != nil && head != nil {
		if slowPtr.Val != head.Val {
			return false
		}

		slowPtr = slowPtr.Next
		head = head.Next
	}

	return true

}

func reverse(head *ListNode) *ListNode {
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
