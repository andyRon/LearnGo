package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Go程序能够智能地处理Unix信号。例如，我们可能希望服务器在收到 SIGTERM 时正常关闭，或者命令行工具在收到 SIGINT 时停止处理输入。

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
go run ch81-signals.go
awaiting signal
^C
interrupt
exiting
*/
