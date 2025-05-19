package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := []byte("界")

	fmt.Println(utf8.Valid(word[:2]))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))

	fmt.Println(utf8.RuneLen('界'))

	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界"))

	p := make([]byte, 3)
	utf8.EncodeRune(p, '好')
	fmt.Println(p)
	fmt.Println(utf8.DecodeRune(p))
	fmt.Println(utf8.DecodeRuneInString("你好"))
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))

	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.FullRuneInString("你好"))

	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}
