package main

import (
	"log"
	"net"
)

/*
goroutine使用模式: 每个请求创建一个goroutine。
这是最常见的模式，每当有新的请求到达时，就创建一个新的goroutine来处理该请求。这样可以充分利用多核CPU的优势，提高服务器的吞吐量。

*/

func main() {
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
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// 处理请求的逻辑
}
