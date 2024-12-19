package main

import "unsafe"

type foo struct {
	id   string
	age  int8
	addr string
}

func main() {
	var p1 *int
	var p2 *bool
	var p3 *byte
	var p4 *[20]int
	var p5 *foo
	var p6 unsafe.Pointer
	println(unsafe.Sizeof(p1)) // 8
	println(unsafe.Sizeof(p2)) // 8
	println(unsafe.Sizeof(p3)) // 8
	println(unsafe.Sizeof(p4)) // 8
	println(unsafe.Sizeof(p5)) // 8
	println(unsafe.Sizeof(p6)) // 8

	var a int = 17
	var p *int = &a
	println(p)
	var p7 *int = &a
	(*p) += 5
	println(*p7) // 22
}
