package main

import (
	"errors"
	"fmt"
)

/*
goroutine使用模式: 错误处理模式
在 goroutine 中执行可能出错的任务时，需要将错误信息传递给主 goroutine 进行处理。可以使用通道传递错误信息。
*/

func taskWithError() <-chan error {
	errChan := make(chan error)
	go func() {
		// 模拟任务出错
		err := errors.New("task failed")
		errChan <- err
		close(errChan)
	}()
	return errChan
}

func main() {
	errChan := taskWithError()
	if err := <-errChan; err != nil {
		fmt.Println("Error:", err)
	}
}
