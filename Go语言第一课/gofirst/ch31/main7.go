package main

import (
	"fmt"
	"time"
)

/*
goroutine使用模式: 生产者 - 消费者模式

生产者 goroutine 负责生成数据，消费者 goroutine 负责处理数据，通过通道进行数据传递。
*/

// 生产者函数
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // 关闭通道，表示数据生产结束
}

// 消费者函数
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consumed: %d\n", num)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch) // 启动生产者 goroutine
	consumer(ch)    // 启动消费者
}
