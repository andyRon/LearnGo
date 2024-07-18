package main

import "fmt"

// 可变参数

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 5)
	sum(1, 5, 8, 7)

	nums := []int{3, 2, 1}
	sum(nums...)
}
