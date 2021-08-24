package remove_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		nums      []int
		search    int
		expectedN int
		expected  []int
	}{
		{
			nums:      []int{},
			search:    1,
			expectedN: 0,
			expected:  []int{},
		},
		{
			nums:      []int{3, 2, 2, 3},
			search:    3,
			expectedN: 2,
			expected:  []int{2, 2, 2, 3},
		},
		{
			nums:      []int{0, 1, 2, 2, 3, 0, 4, 2},
			search:    2,
			expectedN: 5,
			expected:  []int{0, 1, 3, 0, 4, 0, 4, 2},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual := removeElement(c.nums, c.search)
			assert.Equal(t, c.expectedN, actual)
			assert.Equal(t, c.expected, c.nums)
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement([]int{3, 2, 2, 3}, 3)
	}
}
