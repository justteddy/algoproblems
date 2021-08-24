package find_numbers_with_even_number_of_digits

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestFindNumbers(t *testing.T) {
    cases := []struct {
        nums     []int
        expected int
    }{
        {
           nums:     []int{12,345,2,6,7896},
           expected: 2,
        },
        {
           nums:     []int{555,901,482,1771},
           expected: 1,
        },
        {
            nums:     []int{252},
            expected: 0,
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
            assert.Equal(t, c.expected, findNumbers(c.nums))
        })
    }
}

func BenchmarkFindNumbers(b *testing.B) {
    for i := 0; i < b.N; i++ {
        findNumbers([]int{12,345,2,6,7896})
    }
}