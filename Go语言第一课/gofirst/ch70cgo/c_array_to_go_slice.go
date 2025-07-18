package main

// int cArray[] = {1, 2, 3, 4, 5, 6, 7, 8};
import "C"
import (
	"fmt"
	"unsafe"
)

// 将C整型数组转换为Go[]int32切片类型
func main() {
	goArray := CArrayToGoArray(unsafe.Pointer(&C.cArray[0]), unsafe.Sizeof(C.cArray[0]), 8)
	fmt.Println(goArray) // [1 2 3 4 5 6 7 8]
}

func CArrayToGoArray(cArray unsafe.Pointer, elemSize uintptr, len int) (goArray []int32) {
	for i := 0; i < len; i++ {
		j := *(*int32)((unsafe.Pointer)(uintptr(cArray) + uintptr(i)*elemSize))
		goArray = append(goArray, j)
	}
	return
}
