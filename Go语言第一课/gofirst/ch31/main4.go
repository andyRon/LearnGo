package main

import (
	"log"
	"net"
)

/*
goroutine使用模式: 工作池（worker pool）
创建一个固定数量的goroutine作为工作协程，它们会从任务队列中取出任务并执行。这样可以控制并发数，避免过多的goroutine导致系统资源耗尽。
*/

func main() {
	tasks := make(chan func(), 100) // 任务队列，容量为100
	numWorkers := 5                 // 工作协程数量

	for i := 0; i < numWorkers; i++ {
		go newWorker(i, tasks).start()
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		tasks <- func() {
			handleRequest(con)
		}
	}
}

type worker struct {
	id    int
	tasks chan func()
}

func newWorker(id int, tasks chan func()) *worker {
	return &worker{
		id:    id,
		tasks: tasks,
	}
}

func (w *worker) start() {
	for task := range w.tasks {
		task()
	}
}
