package bench

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	tls "github.com/huandu/go-tls"
)

var (
	m     map[int64]struct{} = make(map[int64]struct{}, 10)
	mu    sync.Mutex
	round int64 = 1
)

func BenchmarkSequential(b *testing.B) {
	fmt.Printf("\ngoroutine[%d] enter BenchmarkSequential: round[%d], b.N[%d]\n",
		tls.ID(), atomic.LoadInt64(&round), b.N)
	defer func() {
		atomic.AddInt64(&round, 1)
	}()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		_, ok := m[round]
		if !ok {
			m[round] = struct{}{}
			fmt.Printf("goroutine[%d] enter loop in BenchmarkSequential: round[%d], b.N[%d]\n",
				tls.ID(), atomic.LoadInt64(&round), b.N)
		}
		mu.Unlock()
	}
	fmt.Printf("goroutine[%d] exit BenchmarkSequential: round[%d], b.N[%d]\n",
		tls.ID(), atomic.LoadInt64(&round), b.N)
}

/*
// TODO  go-tls库的问题，可能在是MacOS系统问题
go test -bench . sequential_test.go
panic: tls: fail to call mprotect(addr=0x102c43bb8, size=40, prot=0x2) with error invalid argument

goroutine 1 [running]:
github.com/huandu/go-tls.mprotect(0x102af4e04?, 0x28, 0x2)
        /Users/andyron/myfield/go/pkg/mod/github.com/huandu/go-tls@v1.0.0/syscall_unix.go:34 +0xfc
github.com/huandu/go-tls.init.1()
        /Users/andyron/myfield/go/pkg/mod/github.com/huandu/go-tls@v1.0.0/goexit.go:62 +0x74
exit status 2
FAIL    command-line-arguments  0.423s
FAIL


*/
