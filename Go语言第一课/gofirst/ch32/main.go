package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadloop() {
	runtime.LockOSThread() // 将当前 goroutine 锁定到操作系统线程

	for {
		//fmt.Println("deadloop")
	}
}

func main() {
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

/*
1 在一个拥有多核处理器的主机上，使用Go 1.13.x版本运行这个示例代码，你在命令行终端上是否能看到“I got scheduled!”输出呢？
也就是main goroutine在创建deadloop goroutine之后是否能继续得到调度呢？

能，会
*/

/*
2 通过什么方法可以让上面示例中的main goroutine，在创建deadloop goroutine之后无法继续得到调度？
deadloop 函数包含一个无限空循环，main 函数启动 deadloop 的 goroutine 后，自身也进入一个无限循环，每秒打印一次信息。
正常情况下，Go 的调度器会在多个 goroutine 间切换，所以 main 的 goroutine 能继续调度。
若要让 main 的 goroutine 在创建 deadloop 的 goroutine 后无法继续调度，可从调度器的角度出发，利用一些手段阻塞调度器。

方法一：使用 runtime.LockOSThread 结合死循环
runtime.LockOSThread 会将当前 goroutine 锁定到一个操作系统线程上，若该 goroutine 进入死循环，调度器就没机会切换到其他 goroutine。
func deadloop() {
	runtime.LockOSThread() // 将当前 goroutine 锁定到操作系统线程
	for {
		// 死循环
	}
}

方法二：使用 runtime.LockOSThread 并消耗所有逻辑处理器
通过 runtime.GOMAXPROCS 设置逻辑处理器数量，然后让每个逻辑处理器都运行一个死循环的 goroutine，这样调度器就没资源调度 main 的 goroutine 了。
func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU) // 设置逻辑处理器数量

	for i := 0; i < numCPU; i++ {
		go func() {
			runtime.LockOSThread()
			deadloop()
		}()
	}

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
*/
