package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	nums := []int{1, 2, 3}
	for i := range nums {
		fmt.Println("range", i)
	}
	// 不带条件的 `for` 循环将一直重复执行，直到在循环体内使用 break 或 return
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 9; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
