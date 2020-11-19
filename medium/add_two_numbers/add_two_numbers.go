package add_two_numbers

/*
   You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
   You may assume the two numbers do not contain any leading zero, except the number 0 itself.

   Example 1:
   Input: l1 = [2,4,3], l2 = [5,6,4]
   Output: [7,0,8]
   Explanation: 342 + 465 = 807.

   Example 2:
   Input: l1 = [0], l2 = [0]
   Output: [0]

   Example 3:
   Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
   Output: [8,9,9,9,0,0,0,1]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	resultPtr := dummyHead

	carry := 0
	for l1 != nil || l2 != nil {
		x1 := 0
		if l1 != nil {
			x1 = l1.Val
			l1 = l1.Next
		}
		x2 := 0
		if l2 != nil {
			x2 = l2.Val
			l2 = l2.Next
		}
		sum := x1 + x2 + carry
		carry = sum / 10
		resultPtr.Next = &ListNode{Val: sum % 10}
		resultPtr = resultPtr.Next
	}
	if carry > 0 {
		resultPtr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
