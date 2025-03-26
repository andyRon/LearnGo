package main

import (
	"errors"
	"fmt"
)

// 变量遮蔽

var a int = 2023

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a { // 遮蔽包代码块中的变量(a); 遮蔽外层显式代码块中的变量(err)
	case 2023:
		fmt.Println("it is", a, err)
	case 2024:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)
	return err
}

type new int // 遮蔽预定义标识符

func getYear() (new, error) {
	var b int16 = 2024
	return new(b), nil
}

func main() {
	err := checkYear()
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
