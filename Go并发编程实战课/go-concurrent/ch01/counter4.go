package main

import (
	"fmt"
	"sync"
)

func main() {
	// 封装好的计数器
	var counter Counter2

	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 线程安全的计数器类型
type Counter2 struct {
	CounterType int
	name        string

	mu    sync.Mutex // 如果嵌入的struct有多个字段，我们一般会把Mutex放在要控制的字段上面，然后使用空格把字段分隔开来。
	count uint64
}

// 加1的方法内部使用互拆锁保护
func (c *Counter2) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要互斥锁保护
func (c *Counter2) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
