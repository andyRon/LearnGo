package main

import (
	"fmt"
	"time"
)

/*
goroutine使用模式: 超时控制模式
在调用 goroutine 执行任务时，为避免任务执行时间过长，可以设置超时控制。使用 select 语句和 time.After 函数实现。
*/

func longRunningTask() <-chan string {
	result := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		result <- "Task completed"
	}()
	return result
}

func main() {
	select {
	case res := <-longRunningTask():
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("Task timed out")
	}
}

// time.After 函数会在指定的时长后，向返回的通道发送一个当前时间的 time.Time 值。
// 这个函数主要用于在 select 语句里实现超时逻辑，当超过设定时间后，select 语句就能选择对应的分支执行。

/*
longRunningTask 函数启动一个 goroutine 执行耗时任务，该任务会休眠 3 秒后才向通道发送结果。
main 函数里的 select 语句有两个 case 分支：
	第一个分支尝试从 longRunningTask 返回的通道接收结果。如果在 2 秒内没有接收到结果，会执行第二个分支。
	第二个分支使用 time.After(2 * time.Second)，表示 2 秒后该通道会有数据可读。

由于 longRunningTask 的任务需要 3 秒才能完成，而 time.After 在 2 秒后就会触发，所以 select 语句会执行 Task timed out 这个分支。
*/