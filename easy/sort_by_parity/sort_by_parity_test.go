package sort_by_parity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_by_parity(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{3, 1, 2, 4},
			expected: []int{2, 4, 3, 1},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, sortArrayByParity(c.nums))
		})
	}
}

func BenchmarkSort_by_parity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity([]int{})
	}
}
