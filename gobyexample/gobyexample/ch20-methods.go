package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // 指针接收器
	r.width = 20
	return r.width * r.height
}

func (r rect) perim() int { // 非指针接收器
	r.height = 10
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim()) // 20*2 + 10*2 = 60  指针接收器把width改成了20

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	r2 := rect{width: 10, height: 5}
	fmt.Println("perim:", r2.perim()) // 2*10 + 2*10 = 40
	fmt.Println("area: ", r2.area())  // 20*5 = 100  非指针接收器修改了height但不影响其它

}
