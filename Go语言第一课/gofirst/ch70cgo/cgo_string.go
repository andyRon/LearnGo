package main

// #include <stdio.h>
// #include <stdlib.h>
// char *foo = "helloChina";
import "C"
import "fmt"

func main() {
	fmt.Printf("%T", C.GoString(C.foo)) // string
}
