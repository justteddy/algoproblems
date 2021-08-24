package find_disappeared_numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, c.expected, findDisappearedNumbers(c.nums))
		})
	}
}

func BenchmarkFindDisappearedNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	}
}
