package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 1}

		assert.True(t, isPalindrome(&head))
	})

	t.Run("case 2", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next.Next = &ListNode{Val: 1}

		assert.True(t, isPalindrome(&head))
	})

	t.Run("case 2", func(t *testing.T) {
		head := ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next = &ListNode{Val: 3}
		head.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next = &ListNode{Val: 1}

		assert.False(t, isPalindrome(&head))
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
