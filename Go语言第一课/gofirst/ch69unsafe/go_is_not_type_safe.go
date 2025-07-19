package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint32 = 0x12345678
	fmt.Printf("0x%x\n", a)

	p := (unsafe.Pointer)(&a) // 利用unsafe.Pointer的性质1

	b := (*[4]byte)(p) // 利用unsafe.Pointer的性质2
	b[0] = 0x23
	b[1] = 0x45
	b[2] = 0x67
	b[3] = 0x8a

	fmt.Printf("0x%x\n", a) // 0x8a674523 (注：在小端字节序系统中输出此值)
}
