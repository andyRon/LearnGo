package main

import (
	"fmt"
)

func main() {
	var a int = 0x12345678
	var pa *int = &a
	var pb *byte = (*byte)(pa) // 编译器报错：cannot convert pa (variable of type *int) to type *byte
	fmt.Printf("%x\n", *pb)
}
