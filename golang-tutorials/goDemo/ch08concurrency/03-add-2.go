package main

import (
	"fmt"
	"sync"
	"time"
)

func addV3(a, b int, doneFunc func()) {
	defer func() {
		doneFunc()
	}()
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10) // 初始化计数器数目为10
	for i := 0; i < 10; i++ {
		go addV3(1, i, wg.Done)
	}
	wg.Wait() // 等待子协程全部执行完毕退出
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
