package main

import "fmt"

func main() {
	//printNilInterface()
	//printEmptyInterface()
	//printNonEmptyInterface()
	printEmptyInterfaceAndNonEmptyInterface()
}

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

/*
(0x0,0x0)
(0x0,0x0)
i = nil: true
err = nil: true
i = err: true
*/

func printEmptyInterface() {
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false

	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // true

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false
}

type T int

func (t T) Error() string {
	return "bad error"
}
func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
}

/**
err1: (0x102e30558,0x0)
err1 = nil: false
err1: (0x102e30578,0x102e0feb8)
err2: (0x102e30578,0x102e0fec0)
err1 = err2: false
err1: (0x102e30578,0x102e0feb8)
err2: (0x102e305d8,0x14000010040)
err1 = err2: false
*/

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err) // true

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err) // false
}
