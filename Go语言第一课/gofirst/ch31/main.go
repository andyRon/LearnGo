package main

import (
	"errors"
	"fmt"
	"time"
)

/*
在main goroutine与子goroutine之间建立了一个元素类型为error的channel，
子goroutine退出时，会将它执行的函数的错误返回值写入这个channel，
main goroutine可以通过读取channel的值来获取子goroutine的退出状态。
*/

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func main() {
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}
