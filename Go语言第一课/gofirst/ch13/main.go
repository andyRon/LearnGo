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
	//str3()

	encodeRune()
	//decodeRune()

	//str4()

	//str5()
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

func str3() {
	char := '\u4e2d'
	fmt.Printf("%c\n", char)
	str := "中\u56fd\u4eba"
	fmt.Println(str)
	str = "\U00004e2d\U000056fd\U00004eba"
	fmt.Println(str)
	str = "\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba"
	fmt.Println(str)
}

// 使用Go在标准库中提供的UTF-8包，对Unicode字符（rune）进行编解码

// rune -> []byte
func encodeRune() {
	var r rune = 0x4e2d
	fmt.Printf("这个unicode字符是：%c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r) // 对rune进行utf8编码
	fmt.Printf("这个字符的utf8描述为：0x%X\n", buf)
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xe4, 0xb8, 0xad}
	r, _ := utf8.DecodeRune(buf) // 对buf进行utf8解码
	fmt.Printf("字节序列解码后的unicode字符是：%s\n", string(r))
}

// 字符串内部表示
func str4() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                            // [h e l l o ]   // 输出底层数组的内容
}
func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

// 字符串操作
func str5() {
	var s = "中国人"
	fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节

	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}

	fmt.Println()
	for i, v := range s {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}

	fmt.Println()
	s = "Rob Pike, "
	s = s + "Robert Griesemer, "
	s += " Ken Thompson"
	fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson

	fmt.Println()
	s = "中国人"
	// string -> []rune
	rs := []rune(s)
	fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]

	// string -> []byte
	bs := []byte(s)
	fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba

	// []rune -> string
	s1 := string(rs)
	fmt.Println(s1) // 中国人

	// []byte -> string
	s2 := string(bs)
	fmt.Println(s2) // 中国人

}
