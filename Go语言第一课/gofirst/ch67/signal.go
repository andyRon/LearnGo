package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Singal!\n")
	})
	wg.Add(1)
	go func() {
		defer wg.Done()
		errChan <- http.ListenAndServe(":8080", nil)
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("web server start ok")
	case err := <-errChan:
		fmt.Printf("web server start failed: %v\n", err)

	}
	wg.Wait()
	fmt.Println("web server shutdown ok")
}

/*
go build -o httpserv signal.go
./httpserv
web server start ok
*/
