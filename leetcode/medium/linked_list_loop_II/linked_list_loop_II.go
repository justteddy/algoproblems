package linked_list_loop_II

/*
   Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
   There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
   Notice that you should not modify the linked list.

   Can you solve it using O(1) (i.e. constant) memory?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity - O(n^2)
// Space complexity - O(1)
// we can do it much faster with hashtable or temp node with list modifying
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prevDistance := 0
	ptr := head.Next
	for ptr != nil {
		currDistance := distance(head, ptr)
		if currDistance <= prevDistance {
			return ptr
		}
		prevDistance = currDistance
		ptr = ptr.Next
	}

	return nil
}

func distance(from, target *ListNode) int {
	d := 0
	curr := from
	for curr != target {
		curr = curr.Next
		d++
	}

	return d
}
