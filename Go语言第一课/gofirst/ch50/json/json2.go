package main

import (
	"encoding/json"
	"fmt"
)

// 生成JSON

type server struct {
	ServerName string
	ServerIP   string
}

type serverslice struct {
	Servers []server
}

func main() {
	var s serverslice
	s.Servers = append(s.Servers, server{ServerName: "Local_Web", ServerIP: "172.0.0.1"})
	s.Servers = append(s.Servers, server{ServerName: "Local_DB", ServerIP: "172.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err: ", err)
	}
	fmt.Println(string(b))
}
