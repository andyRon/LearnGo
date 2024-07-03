package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//str1()
	//str2()

	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                            // [h e l l o ]   // 输出底层数组的内容
}

func str1() {
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9

	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Printf("\n")
}

func str2() {
	var s = "中国人"
	fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

	for _, c := range s {
		fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	}
	fmt.Printf("\n")
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}
