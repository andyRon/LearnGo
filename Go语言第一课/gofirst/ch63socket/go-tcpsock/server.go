package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		// ... ...
		// write to the connection
		//... ...
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// 启动一个新的goroutine处理这个新连接￼
		go handleConn(c)
	}
}
