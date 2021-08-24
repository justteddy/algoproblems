package reverse_string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		str      []byte
		expected []byte
	}{
		{
			str:      []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			str:      []byte("Hannah"),
			expected: []byte("hannaH"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			reverseString(c.str)
			assert.Equal(t, c.expected, c.str)
		})
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString([]byte("hello"))
	}
}
