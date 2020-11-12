package reverse_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 0}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := reverseList(&head)
		actual := traverse(newHead)
		assert.Equal(t, []int{5, 4, 3, 2, 1, 0}, actual)
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
