package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf [512]byte
		n, err := c.Read(buf[:])
		if err != nil {
			return
		}
		// write the data to the connection
		_, err = c.Write(buf[:n])
		if err != nil {
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		go handleConn(conn)
	}
}
