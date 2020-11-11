package linked_list_loop_II

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		head := ListNode{
			Val:  0,
			Next: nil,
		}

		head.Next = &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: &head,
			},
		}

		assert.Equal(t, &head, detectCycle(&head))
	})

	t.Run("case 2", func(t *testing.T) {
		head := ListNode{
			Val: 0,
		}

		head.Next = &head

		assert.Equal(t, &head, detectCycle(&head))
	})
}
