package main

import (
	"testing"
	"fmt"
)

func TestName(t *testing.T) {
	fmt.Println(countPrimes(10))
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPrimes(629238)
	}
}

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	sieve := make([]bool, n)
	res := n - 2
	for i := 2; i*i < n; i++ {
		if !sieve[i] {
			for j := i * i; j < n; j += i {
				if !sieve[j] {
					sieve[j] = true
					res--
				}
			}
		}
	}

	return res
}

func countPrimes2(n int) int {
	if n <= 1 {
		return 0
	}

	res := n - 2
	seen := map[int]bool{}

	for i := 2; i*i < n; i++ {
		if _, ok := seen[i]; ok {
			continue
		}
		for j := i * i; j < n; j += i {
			if _, ok := seen[j]; ok {
				continue
			}

			seen[j] = true
			res--
		}
	}

	return res
}
