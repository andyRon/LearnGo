package main

import (
	"fmt"
	"strings"
	"testing"
)

var sl = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(sl)
	}
}

/*
// -8 表示 GOMAXPROCS 的值，这里代表使用 8 个逻辑 CPU 核心进行测试。
// 第二个数字如25480540表示迭代次数，表示在基准测试过程中，BenchmarkConcatStringByOperator 函数中的测试代码片段总共执行的次数。Go 语言的基准测试框架会自动调整这个次数，以保证测试时间足够长，从而得到稳定的性能数据。
// 第三个数表示每次操作的平均耗时：46.90 ns/op，即每次执行测试代码片段平均消耗的时间，单位为纳秒（ns）。
go test -bench . benchmark_intro_test.go
goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkConcatStringByOperator-8       25480540                46.90 ns/op
BenchmarkConcatStringBySprintf-8         5476678               214.4 ns/op
BenchmarkConcatStringByJoin-8           39735648                29.95 ns/op
PASS
ok      command-line-arguments  5.351s


go test -bench=ByJoin ./benchmark_intro_test.go
goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkConcatStringByJoin-8           39560677                30.93 ns/op
PASS
ok      command-line-arguments  2.383s

// 传入-benchmem命令行参数输出内存分配信息（与基准测试代码中显式调用b.ReportAllocs的效果是等价的）
go test -bench=ByJoin ./benchmark_intro_test.go -benchmem
goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkConcatStringByJoin-8           39435440                30.07 ns/op           48 B/op          1 allocs/op
PASS
ok      command-line-arguments  2.191s

*/
