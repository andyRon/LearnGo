package main

import (
	"calc/simplemath"
	"fmt"
	"os"
	"strconv"
)

// 定义一个用于打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

func main() {
	/*
	 * 用于获取命令行参数，注意程序名本身是第一个参数，
	 * 比如 calc add 1 2 这条指令，第一个参数是 calc
	 */
	args := os.Args
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
