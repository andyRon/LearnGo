package main

import (
	"fmt"
	"time"
)

// 无缓冲channel的惯用法1: 信号传递(1对1通知信号)

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
