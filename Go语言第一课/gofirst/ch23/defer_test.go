package main

import "testing"

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}

	return total
}

func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}
func fooWithoutDefer() {
	sum(10)
}

// 测量带有defer的函数执行的性能
func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

// 测量不带有defer的函数的执行的性能
func BenchmarkFooWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}
}

// go test -bench . defer_test.go
