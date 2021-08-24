package majority_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElementTest(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, majorityElement(c.nums))
		})
	}
}

func BenchmarkMajorityElementTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElement([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5})
	}
}
