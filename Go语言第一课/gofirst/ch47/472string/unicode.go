package main

import (
	"fmt"
	"unicode"
)

func main() {
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single))

	digit := '1'
	fmt.Println(unicode.IsDigit(digit))
	fmt.Println(unicode.IsNumber(digit))

	letter := 'Ⅷ'
	fmt.Println(unicode.IsDigit(letter))
	fmt.Println(unicode.IsNumber(letter))

	han := '你'
	fmt.Println(unicode.IsDigit(han))
	fmt.Println(unicode.Is(unicode.Han, han))
	fmt.Println(unicode.In(han, unicode.Gujarati, unicode.White_Space))
}
