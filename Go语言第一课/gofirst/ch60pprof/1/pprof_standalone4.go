package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 独立程序性能数据采集4: 如果是非HTTP服务程序，则在导入包的同时还需单独启动一个用于性能数据采集的goroutine

func main() {
	go func() {
		// 单独启动一个HTTP server用于性能数据采集
		fmt.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-c:
				wg.Done()
				return
			default:
				s1 := "hello,"
				s2 := "gopher"
				s3 := "!"
				_ = s1 + s2 + s3
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("program exit")
}
