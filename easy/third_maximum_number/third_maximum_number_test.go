package third_maximum_number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMaximumNumber(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			nums:     []int{2, 2, 3, 1},
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, thirdMax(c.nums))
		})
	}
}

func BenchmarkThirdMaximumNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		thirdMax([]int{2, 2, 3, 1})
	}
}
