package main

import "sync"

/*
goroutine使用模式：WaitGroup等待退出
场景：主Goroutine需等待一组子任务全部完成（如批量API请求、并行计算）。
注意：Add()必须在go前调用，避免竞态；defer wg.Done()确保异常时仍释放资源。
*/

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // 确保任务完成时计数器减1
			// 执行任务
		}(i)
	}
	wg.Wait() // 阻塞至所有任务完成
}
