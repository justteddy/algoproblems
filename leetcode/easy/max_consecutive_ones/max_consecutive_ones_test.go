package max_consecutive_ones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{0, 1, 0, 1, 1, 1, 0, 1, 1},
			expected: 3,
		},
		{
			nums:     []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			nums:     []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 10,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, findMaxConsecutiveOnes(c.nums))
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMaxConsecutiveOnes([]int{0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1})
	}
}
