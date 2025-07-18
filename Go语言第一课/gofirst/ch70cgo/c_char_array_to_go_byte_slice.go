package main

// char cArray[] = {'a', 'b', 'c', 'd', 'e', 'f'};
import "C"
import (
	"fmt"
	"unsafe"
)

// C中的char类型数组转换为Go中的[]byte切片类型
func main() {
	goArray := C.GoBytes(unsafe.Pointer(&C.cArray[0]), 6)
	fmt.Printf("%c\n", goArray) // [a b c d e f]
}
