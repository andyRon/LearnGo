package main

// #include <stdio.h>
// #include <stdlib.h>
//
// void print(char* s) {
// 	printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	s := "Hello, Cgo"
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.print(cs)
}
