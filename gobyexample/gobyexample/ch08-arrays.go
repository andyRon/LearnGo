package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	b = [...]int{5, 4, 3, 2, 1}
	fmt.Println("dcl: ", b)

	// 如果将索引指定为 : ，则中间的元素将归零。
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx", b)

	var two2D [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two2D[i][j] = i + j
		}
	}
	fmt.Println("2d: ", two2D)

	two2D = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", two2D)
}
