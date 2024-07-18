package main

import "fmt"

// 使用 recover从panic中恢复。 recover 可以阻止 panic 中止程序，并让它继续执行。
// 一个有用的例子：如果一个客户端连接出现严重错误，服务器不想崩溃。相反，服务器会希望关闭该连接并继续为其他客户端提供服务。

func mayPanic() {
	panic("a problem") // TODO
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
