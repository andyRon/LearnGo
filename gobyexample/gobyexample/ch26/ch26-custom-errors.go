package main

import (
	"errors"
	"fmt"
)

// 通过在自定义类型上实现 Error() 方法，可以将自定义类型用作 error 。

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) { // errors.As 是 errors.Is 的更高级版本。它检查给定的错误（或其链中的任何错误）是否匹配特定的错误类型，并转换为该类型的值，返回 true 。如果没有匹配，则返回 false 。
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
