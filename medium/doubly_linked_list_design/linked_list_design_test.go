package doubly_linked_list_design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(0)
		list.AddAtIndex(1, 1)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, -1, list.Get(2))
	})

	t.Run("case 2", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(0)
		list.AddAtTail(1)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
	})

	t.Run("case 3", func(t *testing.T) {
		list := Constructor()
		list.AddAtTail(1)

		assert.Equal(t, 1, list.Get(0))
	})

	t.Run("case 4", func(t *testing.T) {
		list := Constructor()
		list.AddAtTail(0)
		list.AddAtTail(1)
		list.AddAtTail(2)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, 2, list.Get(2))
	})

	t.Run("case 5", func(t *testing.T) {
		list := Constructor()
		list.AddAtTail(0)
		list.AddAtTail(1)
		list.AddAtTail(2)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, 2, list.Get(2))
	})

	t.Run("case 6", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(2)
		list.AddAtHead(1)
		list.AddAtHead(0)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, 2, list.Get(2))
	})

	t.Run("case 7", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(0)
		list.AddAtIndex(1, 2)
		list.AddAtIndex(1, 1)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, 2, list.Get(2))
	})

	t.Run("case 7", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(0)
		list.AddAtIndex(1, 1)
		list.AddAtIndex(2, 2)
		list.AddAtIndex(3, 3)
		list.DeleteAtIndex(3)
		list.DeleteAtIndex(4)

		assert.Equal(t, 0, list.Get(0))
		assert.Equal(t, 1, list.Get(1))
		assert.Equal(t, 2, list.Get(2))
		assert.Equal(t, -1, list.Get(3))
	})

	t.Run("case 8", func(t *testing.T) {
		list := Constructor()
		list.AddAtHead(1)
		list.DeleteAtIndex(0)

		assert.Equal(t, -1, list.Get(0))
	})
}
