package main

import (
	"fmt"
)

func main() {
	a := 0x12345678
	fmt.Printf("%x\n", a)

	var p *byte = (*byte)(&a) // 错误！不允许将&a从*int类型显式转型为*byte类型

	*p = 0x23
	var b byte = byte(a)    // b是一个新变量，有自己所解释的内存空间
	b = 0x23                // 即便强制进行类型转换，原变量a所解释的内存空间的数据依然不变
	fmt.Printf("0x%x\n", b) // 0x23
	fmt.Printf("0x%x\n", a) // 0x12345678￼
}
