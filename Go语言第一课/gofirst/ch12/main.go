package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//int1()

	//print()

	float2()
}

func int1() {
	var a, b = int(5), uint(6)
	var p uintptr = 0x12345678
	fmt.Println("signed integer a's length is", unsafe.Sizeof(a))   // 8
	fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
	fmt.Println("uintptr's length is", unsafe.Sizeof(p))            // 8
}

func print() {
	var a int8 = 59
	fmt.Printf("%b\n", a) //输出二进制：111011
	fmt.Printf("%d\n", a) //输出十进制：59
	fmt.Printf("%o\n", a) //输出八进制：73
	fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
	fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
	fmt.Printf("%X\n", a) //输出十六进制(大写)：3B
}

func float() {
	var f float32 = 139.8125
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits) // 1000011000010111101000000000000
}

func float2() {
	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	fmt.Println(f1 == f2) // true
}
