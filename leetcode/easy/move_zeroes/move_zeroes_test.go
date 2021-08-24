package move_zeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			moveZeroes(c.nums)
			assert.Equal(t, c.expected, c.nums)
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moveZeroes([]int{0, 1, 0, 3, 12})
	}
}
