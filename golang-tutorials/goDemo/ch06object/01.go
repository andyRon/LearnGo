package main

import "fmt"

type Integer int

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

func main() {
	//var x Integer
	//var y Integer
	//x, y = 10, 11
	//fmt.Println(x.Equal(y))
	//fmt.Println(x.Add(y))
	//fmt.Println(x.Multiply(y))

	var a Integer = 1
	var m Math = a
	fmt.Println(m.Add(12))
}

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}
