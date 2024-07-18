package main

import (
	"fmt"
	"os"
)

// go run ch71-command-line-arguments.go a c e d
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:] // 参数

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
