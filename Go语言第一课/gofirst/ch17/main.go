package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var t T

	fmt.Println(unsafe.Sizeof(t))      // 结构体类型变量占用的内存大小
	fmt.Println(unsafe.Offsetof(t.Fn)) // 字段Fn在内存中相对于变量t起始地址的偏移量
	fmt.Println(unsafe.Sizeof(t.F1.b))
	fmt.Println(unsafe.Sizeof(t.F1.u))

}

type T struct {
	F1 T1
	F2 T2
	F3 T3
	Fn Tn
}

type T1 struct {
	b byte
	i int64
	u uint16
}
type T2 struct {
	b int
}
type T3 struct {
}
type Tn struct {
}
