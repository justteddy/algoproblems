package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		var l1 *ListNode

		l2 := &ListNode{Val: 2}
		l2.Next = &ListNode{Val: 3}

		result := traverse(mergeTwoLists(l1, l2))
		assert.Equal(t, []int{2, 3}, result)
	})

	t.Run("case 2", func(t *testing.T) {
		var l2 *ListNode

		l1 := &ListNode{Val: 2}
		l1.Next = &ListNode{Val: 3}

		result := traverse(mergeTwoLists(l1, l2))
		assert.Equal(t, []int{2, 3}, result)
	})

	t.Run("case 3", func(t *testing.T) {
		l1 := &ListNode{Val: 1}
		l1.Next = &ListNode{Val: 2}
		l1.Next.Next = &ListNode{Val: 3}

		l2 := &ListNode{Val: 2}
		l2.Next = &ListNode{Val: 3}

		result := traverse(mergeTwoLists(l1, l2))
		assert.Equal(t, []int{1, 2, 2, 3, 3}, result)
	})

	t.Run("case 4", func(t *testing.T) {
		l1 := &ListNode{Val: 1}
		l1.Next = &ListNode{Val: 2}
		l1.Next.Next = &ListNode{Val: 4}

		l2 := &ListNode{Val: 1}
		l2.Next = &ListNode{Val: 3}
		l2.Next.Next = &ListNode{Val: 4}

		result := traverse(mergeTwoLists(l1, l2))
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, result)
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
