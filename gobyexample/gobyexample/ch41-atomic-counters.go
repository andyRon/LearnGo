package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子计数器

// 在Go中管理状态的主要机制是通过通道进行通信。
// 也可以使用 sync/atomic 包来处理多个goroutine访问的原子计数器

func main() {

	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
