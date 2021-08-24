package duplicate_zeros

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateZeros(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		//{
		//	nums:     []int{1, 0, 2, 3, 0, 4, 5, 0},
		//	expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		//},
		//{
		//	nums:     []int{1, 2, 3},
		//	expected: []int{1, 2, 3},
		//},
		//{
		//	nums:     []int{0, 1, 7, 6, 0, 2, 0, 7},
		//	expected: []int{0, 0, 1, 7, 6, 0, 0, 2},
		//},
		{
			nums:     []int{8, 4, 5, 0, 0, 0, 0, 7},
			expected: []int{8, 4, 5, 0, 0, 0, 0, 0},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			duplicateZeros(c.nums)
			assert.Equal(t, c.expected, c.nums)
		})
	}
}

func BenchmarkDuplicateZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
	}
}
