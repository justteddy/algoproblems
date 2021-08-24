package linked_list_loop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{
			Val:  1,
			Next: nil,
		}

		head.Next = &ListNode{
			Val:  2,
			Next: nil,
		}

		assert.False(t, hasCycle(&head))
	})
}
