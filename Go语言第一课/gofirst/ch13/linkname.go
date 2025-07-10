package main

import (
	"fmt"
	"unsafe"
)

// 字符串内部表示2，StringHeader过期，使用go:linkname

func main() {
	//str4_2()  // TODO
}

// go:linkname 是 Go 编译器的一个特殊指令，它允许在一个包中引用另一个包的未导出标识符（如函数、变量等）。
// 不过，使用 go:linkname 会绕过 Go 的封装和类型检查机制，可能带来安全风险和兼容性问题，因此要谨慎使用。
// 下面以访问字符串底层结构为例，展示 go:linkname 的使用案例。
//
//go:linkname stringStruct runtime.stringStruct
type stringStruct struct {
	str unsafe.Pointer
	len int
}

// getStringStruct 使用 go:linkname 获取字符串的底层结构
func getStringStruct(s string) *stringStruct {
	return (*stringStruct)(unsafe.Pointer(&s))
}
func str4_2() {
	s := "hello"
	// 获取字符串的底层结构
	strStruct := getStringStruct(s)
	fmt.Printf("底层字节数组地址: %p\n", strStruct.str)
	fmt.Printf("字符串长度: %d\n", strStruct.len)
}
