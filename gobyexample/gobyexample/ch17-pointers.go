package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0 // *iptr表示这个指针所指向的值
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // &i表示这个变量的指针
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
