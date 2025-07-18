package main

import "testing"

func BenchmarkCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallCFunc()
	}
}

func BenchmarkGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallGoFunc()
	}
}
