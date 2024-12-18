package main

type C1 interface {
	~int | ~int32
	M1()
}

type T struct{}

func (T) M1() {
}

type T1 int

func (T1) M1() {
}

func foo[P C1](t P) {

}

func main() {
	var t1 T1
	foo(t1)
	//var t T
	//foo(t) // 编译器报错： Cannot use T as the type C1. Type does not implement constraint 'C1' because type is not included in type set ('~int', '~int32')
}
