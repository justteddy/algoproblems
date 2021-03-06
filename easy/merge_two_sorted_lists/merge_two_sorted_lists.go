package merge_two_sorted_lists

/*
   Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

    Example 1:
    Input: l1 = [1,2,4], l2 = [1,3,4]
    Output: [1,1,2,3,4,4]

    Example 2:
    Input: l1 = [], l2 = []
    Output: []

    Example 3:
    Input: l1 = [], l2 = [0]
    Output: [0]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	result := &ListNode{}
	resultPtr := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			resultPtr.Next = l1
			l1 = l1.Next
		} else {
			resultPtr.Next = l2
			l2 = l2.Next
		}
		resultPtr = resultPtr.Next
	}

	if l1 == nil {
		resultPtr.Next = l2
	}

	if l2 == nil {
		resultPtr.Next = l1
	}

	return result.Next

}
