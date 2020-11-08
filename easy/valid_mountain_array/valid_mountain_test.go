package valid_mountain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMountain(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 3, 2},
			expected: true,
		},
		{
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: false,
		},
		{
			nums:     []int{2, 1},
			expected: false,
		},
		{
			nums:     []int{3, 5, 5},
			expected: false,
		},
		{
			nums:     []int{0, 3, 2, 1},
			expected: true,
		},
		{
			nums:     []int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11},
			expected: true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, validMountainArray(c.nums))
		})
	}
}

func BenchmarkValidMountain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validMountainArray([]int{0, 3, 2, 1})
	}
}
