package main

import "testing"

func BenchmarkGenerateLargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateLargeString(1000)
	}
}