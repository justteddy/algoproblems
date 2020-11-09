package height_checker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeightChecker(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
		{
			nums:     []int{5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, heightChecker(c.nums))
		})
	}
}

func BenchmarkHeightChecker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heightChecker([]int{1, 1, 4, 2, 1, 3})
	}
}
