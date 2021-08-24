package linked_list_intersection

/*
   Write a program to find the node at which the intersection of two singly linked lists begins.
    Notes:
        If the two linked lists have no intersection at all, return null.
        The linked lists must retain their original structure after the function returns.
        You may assume there are no cycles anywhere in the entire linked structure.
        Each value on each linked list is in the range [1, 10^9].
        Your code should preferably run in O(n) time and use only O(1) memory.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dict := make(map[*ListNode]struct{})
	traverseWithCheck(headA, dict)
	return traverseWithCheck(headB, dict)
}

func traverseWithCheck(head *ListNode, dict map[*ListNode]struct{}) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		if _, ok := dict[curr]; ok {
			return curr
		}
		dict[curr] = struct{}{}
		curr = curr.Next
	}
	return nil
}
