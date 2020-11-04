package single_number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual := singleNumber(c.nums)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber([]int{4, 1, 2, 1, 2, 1})
	}
}
