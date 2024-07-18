package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe()) // 由于 container 嵌入了 base ， base 的方法也变成了 container 的方法

	type describer interface {
		describe() string
	}

	var d describer = co // base实现了describe方法，container也就实现了这个方法，而实现了describe方法，就等于实现了describer接口
	fmt.Println("describer:", d.describe())
}
