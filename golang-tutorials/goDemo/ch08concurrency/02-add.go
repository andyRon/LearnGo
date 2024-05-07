package main

import (
	"fmt"
	"time"
)

func add(a, b int) {
	var c = a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	for i := 0; i < 10; i++ {
		go add(1, i)
	}
	time.Sleep(1 * time.Second) // 让主协程等待1s后退出
}
