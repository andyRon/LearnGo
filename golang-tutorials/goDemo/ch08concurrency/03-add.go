package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func addV2(a, b int, lock *sync.Mutex) {
	lock.Lock()
	c := a + b
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{} // 互斥锁
	for i := 0; i < 10; i++ {
		go addV2(1, i, lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() // 让出CPU的时间片
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时（s): ", consume)
}
