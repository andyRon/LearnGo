package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串和rune类型

// Go语言中的字符串是一个只读的byte类型的切片。
//Go语言和标准库特别对待字符串 - 作为以 UTF-8 为编码的文本容器。
//在其他语言当中， 字符串由”字符”组成。 在Go语言当中，字符的概念被称为 rune - 它是一个表示 Unicode 编码的整数。

func main() {

	const s = "สวัสดี" // 表示泰语中单词“hello”
	// go字符串等价于 []byte
	// 泰语"สวัสดี"有六个字符
	// utf8泰语的每个字符是3个字节
	// len是求字节数量，3 * 6
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ { // 对字符串进行索引会在每个索引处生成原始字节值。 这个循环生成构成s中 Unicode 的所有字节的十六进制值。
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println()

	// 要计算字符串中有多少rune，我们可以使用utf8包。
	// 注意RuneCountInString的运行时取决于字符串的大小。 因为它必须按顺序解码每个 UTF-8 rune。
	// 一些泰语字符由多个 UTF-8 code point 表示， 所以这个计数的结果可能会令人惊讶。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串并解码每个 rune 及其在字符串中的偏移量。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// 通过显式使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// 单引号括起来的值是 rune literals. 可以直接将 rune value 与 rune literal 进行比较。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
