package main

type E interface {
	M1()
	M2()
}

type I interface {
	M1()
	M2()
	M3()
}

type I2 interface {
	E
	M3()
}
