package remove_duplicates_sorted_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums      []int
		expectedN int
		expected  []int
	}{
		{
			nums:      []int{1, 1, 2},
			expectedN: 2,
			expected:  []int{1, 2, 2},
		},
		{
			nums:      []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedN: 5,
			expected:  []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual := removeDuplicates(c.nums)
			assert.Equal(t, c.expectedN, actual)
			assert.Equal(t, c.expected, c.nums)
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}
