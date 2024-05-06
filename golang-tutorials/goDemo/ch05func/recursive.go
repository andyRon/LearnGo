package main

import (
	"fmt"
	"time"
)

type FibonacciFunc func(int) int

// 通过递归函数实现斐波那契数列
func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 斐波那契函数执行耗时计算
func fibonacciExecTime(f FibonacciFunc) FibonacciFunc {
	return func(n int) int {
		start := time.Now()
		num := f(n)
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("-----执行耗时：%v ------\n", end)
		return num
	}
}

const MAX = 50

var fibs [MAX]int

// 缓存中间结果的递归函数优化版
func fibonacci2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fibonacci2(n-1) + fibonacci2(n-2)
	fibs[index] = num
	return num
}

// 尾递归
func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second)
}
func fibonacci3(n int) int {
	return fibonacciTail(n, 0, 1) // F(1) = 0, F(2) = 1
}

func main() {
	//n := 5
	//num := fibonacci(n)
	//fmt.Printf("The %dth number of fibonacci sequence is %d\n", n, num)

	n1 := 5
	f := fibonacciExecTime(fibonacci)
	r1 := f(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)
	n2 := 50
	r2 := f(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)

	f2 := fibonacciExecTime(fibonacci2)
	r3 := f2(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r3)

	f3 := fibonacciExecTime(fibonacci3)
	r4 := f3(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r4)
}
