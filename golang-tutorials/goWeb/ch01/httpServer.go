package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// r 代表请求对象，w 代表响应对象
func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // 在服务端打印请求参数
	fmt.Println("URL: ", r.URL.Path)
	fmt.Println("Scheme: ", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好，这是andy的go服务器") // 发送响应到客服端
}

func main() {
	http.HandleFunc("/", sayHelloWorld)      // 定义了一个路由`/`和对应的路由处理函数  TODO
	err := http.ListenAndServe(":9091", nil) // 监听端口9091
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
