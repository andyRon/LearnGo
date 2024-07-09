package main

import "unsafe"

// 不同的字段排列顺序也会影响到“填充字节”的多少，从而影响到整个结构体大小

type T struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

func main() {
	var t T
	println(unsafe.Sizeof(t)) // 24
	var s S
	println(unsafe.Sizeof(s)) // 16
}
