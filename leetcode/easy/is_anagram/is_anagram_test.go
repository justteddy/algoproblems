package is_anagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{
			s1:       "anagram",
			s2:       "nagaram",
			expected: true,
		},
		{
			s1:       "rat",
			s2:       "car",
			expected: false,
		},
		{
			s1:       "aa",
			s2:       "bb",
			expected: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual := isAnagram(c.s1, c.s2)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("anagram", "nagaram")
	}
}
