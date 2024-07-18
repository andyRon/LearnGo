package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p) // 包括结构体的字段名称
	fmt.Printf("struct3: %#v\n", p) // 产生该值的源代码片段
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456) // 十六进制
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	fmt.Printf("str1: %s\n", "\"string\"") // 基本的字符串
	fmt.Printf("str2: %q\n", "\"string\"") // 像在Go源代码中那样使用双引号
	fmt.Printf("str3: %x\n", "hex this")   // %x 以16为基数呈现字符串，每个输入字节有两个输出字符
	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // - 左对齐
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	//  Printf将格式化的字符串打印到 os.Stdout 。 Sprintf 格式化并返回一个字符串，而不打印它
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // Fprintf格式化并打印到 io.Writers 而不是 os.Stdout
}
