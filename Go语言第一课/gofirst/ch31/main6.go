package main

import (
	"context"
	"fmt"
	"time"
)

/*
goroutine使用模式: Context生命周期管理
场景：需超时控制、取消传播或传递请求上下文（如HTTP请求处理）。
优势：支持级联取消，避免goroutine泄漏。
*/

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 监听取消信号
			fmt.Println("Canceled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go worker2(ctx)
	time.Sleep(3 * time.Second) // 超时后自动触发ctx.Done()
}
