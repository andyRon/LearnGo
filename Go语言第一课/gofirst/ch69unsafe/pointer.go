package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uintptr = 0x80010203
	p := unsafe.Pointer(i)
	fmt.Println(p)
}

//func p1() {
//	var a int = 5
//	var b float64 = 5.89
//	var arr [10]string
//	var f Foo
//
//	p1 := (unsafe.Pointer)(&a)   // *int -> unsafe.Pointer
//	p2 := (unsafe.Pointer)(&b)   // *float64 -> unsafe.Pointer
//	p3 := (unsafe.Pointer)(&arr) // *[10]string -> unsafe.Pointer
//	p4 := (unsafe.Pointer)(&f)   // *Foo -> unsafe.Pointer
//
//	var pa = (*int)(p1)          // unsafe.Pointer -> *int
//	var pb = (*float64)(p2)      // unsafe.Pointer -> *float64
//	var parr = (*[10]string)(p3) // unsafe.Pointer -> *[10]string
//	var pf = (*Foo)(p4)          // unsafe.Pointer -> *Foo
//}
//
//func Foo() {}
