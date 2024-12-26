package main

import "fmt"

func main() {
	// 后进先出
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
