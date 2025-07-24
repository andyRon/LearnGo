package main

import (
	"expvar"
	"fmt"
	"net/http"
)

// 手动将expvar包的服务端点注册到应用程序所使用的“路由器”上
func main() {
	mux := http.NewServeMux()
	mux.Handle("/hi", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi2"))
	}))
	mux.Handle("/debug/vars", expvar.Handler())
	fmt.Println(http.ListenAndServe("localhost:8080", mux))
}

// http://localhost:8080/debug/vars
