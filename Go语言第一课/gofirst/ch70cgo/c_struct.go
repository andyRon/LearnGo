package main

// #include <stdlib.h>
//
// struct employee {
//   char *id;
//   int age;
// };
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	id := C.CString("123456")
	defer C.free(unsafe.Pointer(id))

	var p = C.struct_employee{
		id:  id,
		age: 18,
	}
	// %#v 以 Go 语法格式输出值
	fmt.Printf("%#v\n", p) // main._Ctype_struct_employee{id:(*main._Ctype_char)(0x600002efc000), age:18, _:[4]uint8{0x0, 0x0, 0x0, 0x0}}
}
