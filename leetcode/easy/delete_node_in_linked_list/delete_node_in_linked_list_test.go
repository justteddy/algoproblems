package delete_node_in_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNodeInLinkedList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {

		l1 := &ListNode{Val: 1}
		l2 := &ListNode{Val: 2}
		l3 := &ListNode{Val: 3}
		l4 := &ListNode{Val: 4}

		l1.Next = l2
		l2.Next = l3
		l3.Next = l4

		deleteNode(l3)
		result := traverse(l1)
		assert.Equal(t, []int{1, 2, 4}, result)
	})
}

func traverse(head *ListNode) []int {
	result := make([]int, 0)
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}
