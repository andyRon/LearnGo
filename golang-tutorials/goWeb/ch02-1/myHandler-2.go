package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (handler *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sayHelloGolang2(w, r)
}

type WorldHandler struct {
}

func (handler *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func sayHelloGolang2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang!!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: ":9091",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
