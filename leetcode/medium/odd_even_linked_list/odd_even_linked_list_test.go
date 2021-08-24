package odd_even_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenLinkedList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next = &ListNode{Val: 5}
		head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

		newHead := oddEvenList(&head)
		actual := traverse(newHead)
		assert.Equal(t, []int{1, 3, 5, 2, 4, 6}, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := oddEvenList(&head)
		actual := traverse(newHead)
		assert.Equal(t, []int{1, 3, 5, 2, 4}, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}

		newHead := oddEvenList(&head)
		actual := traverse(newHead)
		assert.Equal(t, []int{1, 2}, actual)
	})

	t.Run("case 4", func(t *testing.T) {
		head := ListNode{Val: 1}

		newHead := oddEvenList(&head)
		actual := traverse(newHead)
		assert.Equal(t, []int{1}, actual)
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
