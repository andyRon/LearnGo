package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 行过滤器是一种常见的程序类型，它读取stdin上的输入，处理它，然后将一些派生结果打印到stdout。 grep 和 sed 是公用线路滤波器。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
