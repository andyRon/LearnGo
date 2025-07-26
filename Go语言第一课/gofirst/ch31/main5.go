package main

import (
	"log"
	"net"
)

/*
goroutine使用模式: 缓冲通道（buffered channel）
通过限制通道的容量，可以控制同时处理的任务数量。当通道满时，发送者会被阻塞，直到有空闲空间可用。这可以防止过多的任务堆积在内存中。
*/

const maxConcurrency = 10 // 同时处理的最大任务数量

func main() {
	tasks := make(chan func(), maxConcurrency) // 带有缓冲区的通道
	// 启动工作协程
	for i := 0; i < maxConcurrency; i++ {
		go func() {
			for task := range tasks {
				task()
			}
		}()
	}

	// 监听端口并接收请求
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		tasks <- func() {
			handleRequest(conn)
		}
	}

}
