package merge_sorted_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{0, 1},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			merge(c.nums1, c.m, c.nums2, c.n)
			assert.Equal(t, c.expected, c.nums1)
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}
