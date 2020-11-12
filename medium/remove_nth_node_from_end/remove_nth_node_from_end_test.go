package remove_nth_node_from_end

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := removeNthFromEnd(&head, 2)
		assert.Equal(t, []int{0, 1, 2, 3, 5}, traverse(newHead))
	})

	t.Run("case 3", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		newHead := removeNthFromEnd(&head, 1)

		assert.Equal(t, []int{0}, traverse(newHead))
	})

	t.Run("case 4", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		newHead := removeNthFromEnd(&head, 2)

		assert.Equal(t, []int{1}, traverse(newHead))
	})

	t.Run("case 5", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		newHead := removeNthFromEnd(&head, 1)
		assert.Equal(t, []int{0, 1}, traverse(newHead))
	})

	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}

		newHead := removeNthFromEnd(&head, 2)
		assert.Equal(t, []int{0, 2}, traverse(newHead))
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
