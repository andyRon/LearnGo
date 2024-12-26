package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}

/*
	实现了fmt.Stringer这个interface，即如果需要某个类型能被fmt包以特殊的格式输出，你就必须实现Stringer这个接口。
	如果没有实现这个接口，fmt将以默认的方式输出。
*/
