package check_if_double_exist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfExist(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{4, 4},
			expected: false,
		},
		{
			nums:     []int{4, -7, 11, 4, 18},
			expected: false,
		},
		{
			nums:     []int{10, 2, 5, 3},
			expected: true,
		},
		{
			nums:     []int{7, 1, 14, 11},
			expected: true,
		},
		{
			nums:     []int{3, 1, 7, 11},
			expected: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, checkIfExist(c.nums))
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkIfExist([]int{2, 7, 11, 15, 20, 99, 30, 50, 101, 1002, 7854})
	}
}
