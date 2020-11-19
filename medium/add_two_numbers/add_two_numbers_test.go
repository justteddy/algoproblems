package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		l1 := &ListNode{Val: 4}
		l1.Next = &ListNode{Val: 3}
		l1.Next.Next = &ListNode{Val: 3}

		l2 := &ListNode{Val: 2}
		l2.Next = &ListNode{Val: 3}

		result := traverse(addTwoNumbers(l1, l2))
		assert.Equal(t, []int{6, 6, 3}, result)
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
