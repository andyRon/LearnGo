package main

type T struct {
	n int
	s string
}

func (T) M1() {}
func (T) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T{
		n: 18,
		s: "hello, interface",
	}
	var i NonEmptyInterface = t

	println(i) // (0x1008583f8,0x14000050720)
}
