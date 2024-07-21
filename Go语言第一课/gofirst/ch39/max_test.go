package main

import "testing"

func BenchmarkMaxInt(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxInt(sl)
	}
}

func BenchmarkMaxAny(b *testing.B) {
	sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxAny(sl)
	}
}
