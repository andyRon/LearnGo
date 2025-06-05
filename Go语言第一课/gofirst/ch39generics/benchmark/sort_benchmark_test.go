package main

import (
	"golang.org/x/exp/slices"
	"math/rand"
	"sort"
	"testing"
)

const N = 100_000

// https://github.com/golang/exp/blob/master/slices/sort_benchmark_test.go

func makeRandomInts(n int) []int {
	rand.Seed(42)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(n)
	}
	return ints
}

// Go标准库sort包（非泛型版）的Ints函数
func BenchmarkSortInts(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeRandomInts(N)
		b.StartTimer()
		sort.Ints(ints)
	}
}

// Go团队维护golang.org/x/exp/slices中的泛型版Sort函数；
func BenchmarkSlicesSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeRandomInts(N)
		b.StartTimer()
		slices.Sort(ints)
	}
}

// 对golang.org/x/exp/slices中的泛型版Sort函数进行改造得到的、仅针对[]int进行排序的Sort函数。
func BenchmarkIntSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeRandomInts(N)
		b.StartTimer()
		Sort(ints)
	}
}
