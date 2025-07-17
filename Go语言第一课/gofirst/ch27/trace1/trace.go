package main

import "runtime"

func main() {
	defer Trace()()
	foo()
}

func Trace() func() {
	// runtime.Caller函数获得当前Goroutine的函数调用栈上的信息
	// runtime.Caller的参数标识的是要获取的是哪一个栈帧的信息
	// 0表示返回的是Caller函数的调用者的函数信息，在这里就是Trace函数。1是Trace函数的调用者的信息。
	// 四个返回值分别是：程序计数（pc）；函数所在的源文件名以及所在行数；是否能成功获取这些信息；
	pc, fileName, line, ok := runtime.Caller(1)
	println("pc: ", pc, " fileName: ", fileName, " line: ", line, " ok: ", ok)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc) // 被跟踪函数的函数名称
	name := fn.Name()

	println("enter: ", name)
	return func() {
		println("exit: ", name)
	}
}

func foo() {
	defer Trace()()
	bar()
}
func bar() {
	defer Trace()()
}
