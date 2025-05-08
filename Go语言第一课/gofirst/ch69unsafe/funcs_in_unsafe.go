package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Foo struct {
		a int
		b string
		c [10]byte
		d float64
	}
	var i int = 5
	var a = [100]int{}
	var sl = a[:]
	var f Foo
	var s string

	fmt.Println(unsafe.Sizeof(i))           // 8￼
	fmt.Println(unsafe.Sizeof(a))           // 800
	fmt.Println(unsafe.Sizeof(sl))          // 24 (注：返回的是切片描述符的大小)
	fmt.Println(unsafe.Sizeof(f))           // 48
	fmt.Println(unsafe.Sizeof(f.c))         // 10
	fmt.Println(unsafe.Sizeof((*int)(nil))) // 8
	fmt.Println(unsafe.Sizeof(f.b))
	fmt.Println(unsafe.Sizeof(s))

	fmt.Println("----------")

	fmt.Println(unsafe.Alignof(i))          // 8￼
	fmt.Println(unsafe.Alignof(a))          // 8
	fmt.Println(unsafe.Alignof(sl))         // 8
	fmt.Println(unsafe.Alignof(f))          // 8
	fmt.Println(unsafe.Alignof(f.c))        // 1
	fmt.Println(unsafe.Alignof(struct{}{})) // 1  (注：空结构体的对齐系数为1)￼
	fmt.Println(unsafe.Alignof([0]int{}))   // 8 (注：长度为0的数组，其对齐系数依然与其元素类型的对齐系数相同)

	fmt.Println("----------")

	fmt.Println(unsafe.Offsetof(f.a)) // 0
	fmt.Println(unsafe.Offsetof(f.b)) // 8
	fmt.Println(unsafe.Offsetof(f.c)) // 24
	fmt.Println(unsafe.Offsetof(f.d)) // 40
}
