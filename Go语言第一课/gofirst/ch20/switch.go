package main

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}
func switch1() {
	switch switchexpr() {
	case case1():
		println("exec case1")
	case case2_1(), case2_2():
		println("exec case2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}
}

type person struct {
	name string
	age  int
}

func switch2() {
	p := person{"tom", 13}
	switch p {
	case person{"tony", 33}:
		println("match tony")
	case person{"tom", 13}:
		println("match tom")
	case person{"lucy", 23}:
		println("match lucy")
	default:
		println("no match")
	}
}
func main() {
	switch2()
}
