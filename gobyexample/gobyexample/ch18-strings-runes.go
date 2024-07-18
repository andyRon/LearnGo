package main

import (
	"fmt"
	"unicode/utf8"
)

// TODO
func main() {

	const s = "สวัสดี" // 表示泰语中单词“hello”
	// go字符串等价于 []byte
	// 泰语"สวัสดี"有六个字符
	// utf8泰语的每个字符是3个字节
	// len是求字节数量，3 * 6
	fmt.Println("Len:", len(s))

	//
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
