package main

func f(pp **int) {
	var b int = 55
	var p1 *int = &b
	(*pp) = p1
}
func main() {
	var a int = 5
	var p *int = &a
	println(*p) // 5
	f(&p)
	println(*p) // 55
}
