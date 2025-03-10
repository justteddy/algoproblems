package main

import "testing"

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords("   hello     world   ")
	}
}

func BenchmarkReverseWordsStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords2("   hello     world   ")
	}
}
