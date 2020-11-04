package fizz_buzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		n        int
		expected []string
	}{
		{
			n: 15,
			expected: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual := fizzBuzz(c.n)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(15)
	}
}
