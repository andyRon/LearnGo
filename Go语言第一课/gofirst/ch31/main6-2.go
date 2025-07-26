package main

import (
	"context"
	"fmt"
	"time"
)

/*
goroutine使用模式: Context 取消模式
使用 context 包可以在多个 goroutine 之间传递取消信号，实现任务的取消操作。
*/

// workerWithCancel 函数是一个工作协程，会持续执行任务，直到接收到取消信号。main 函数启动该工作协程，并在 2 秒后发送取消信号。

func workerWithCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithCancel(ctx)

	time.Sleep(2 * time.Second)
	cancel() // 取消任务
	time.Sleep(1 * time.Second)
}
