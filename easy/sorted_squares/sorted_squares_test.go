package sorted_squares

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, sortedSquares(c.nums))
		})
	}
}

func BenchmarkSortedSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortedSquares([]int{-7, -3, 2, 3, 11})
	}
}
