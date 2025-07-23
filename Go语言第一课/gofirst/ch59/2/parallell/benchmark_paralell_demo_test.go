package parallell

import (
	"sync"
	"sync/atomic"
	"testing"
)

var n1 int64

func addSyncByAtomic(delta int64) int64 {
	return atomic.AddInt64(&n1, delta)
}

func readSyncByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}

var n2 int64
var rwmu sync.RWMutex

func addSyncByMutex(delta int64) {
	rwmu.Lock()
	n2 += delta
	rwmu.Unlock()
}

func readSyncByMutex() int64 {
	var n int64
	rwmu.RLock()
	n = n2
	rwmu.RUnlock()
	return n
}

func BenchmarkAddSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByAtomic(1)
		}
	})
}

func BenchmarkReadSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByAtomic()
		}
	})
}

func BenchmarkAddSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByMutex(1)
		}
	})
}
func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByMutex()
		}
	})
}

/*
// -cpu 2,4,8命令行选项告知go test将每个性能基准测试函数分别在GOMAXPROCS等于2、4、8的情况下各运行一次。
go test -v -bench . benchmark_paralell_demo_test.go -cpu 2,4,8
goos: darwin
goarch: arm64
cpu: Apple M1
BenchmarkAddSyncByAtomic
BenchmarkAddSyncByAtomic-2      68508310                17.31 ns/op
BenchmarkAddSyncByAtomic-4      54367728                24.22 ns/op
BenchmarkAddSyncByAtomic-8      30219247                40.87 ns/op
BenchmarkReadSyncByAtomic
BenchmarkReadSyncByAtomic-2     1000000000               0.4080 ns/op
BenchmarkReadSyncByAtomic-4     1000000000               0.2228 ns/op
BenchmarkReadSyncByAtomic-8     1000000000               0.1759 ns/op
BenchmarkAddSyncByMutex
BenchmarkAddSyncByMutex-2       23990422                43.14 ns/op
BenchmarkAddSyncByMutex-4       14498713                79.53 ns/op
BenchmarkAddSyncByMutex-8       13313528                93.14 ns/op
BenchmarkReadSyncByMutex
BenchmarkReadSyncByMutex-2      34542397                34.83 ns/op
BenchmarkReadSyncByMutex-4      27798422                43.49 ns/op
BenchmarkReadSyncByMutex-8      15354471                79.86 ns/op
PASS
ok      command-line-arguments  13.402s

*/
