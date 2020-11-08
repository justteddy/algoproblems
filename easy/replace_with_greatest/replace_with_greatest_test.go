package replace_with_greatest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{17, 18, 5, 4, 6, 1},
			expected: []int{18, 6, 6, 6, 1, -1},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, replaceElements(c.nums))
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceElements([]int{})
	}
}
