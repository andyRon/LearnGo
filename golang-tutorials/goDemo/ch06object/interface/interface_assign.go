package main

import "fmt"

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
	Add(i int)
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}

func (n *Number) Add(i int) {
	*n = *n + Number(i)
}

func main() {
	var num1 Number = 1
	var num2 Number2 = &num1
	var num3 Number1 = num2
	fmt.Println(num3)
}
