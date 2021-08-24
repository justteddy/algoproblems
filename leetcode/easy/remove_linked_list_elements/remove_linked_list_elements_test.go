package remove_linked_list_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := removeElements(&head, 3)
		actual := traverse(newHead)
		assert.Equal(t, []int{0, 1, 2, 4, 5}, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		head := ListNode{Val: 2}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := removeElements(&head, 2)
		actual := traverse(newHead)
		assert.Equal(t, []int{3, 4, 5}, actual)
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
