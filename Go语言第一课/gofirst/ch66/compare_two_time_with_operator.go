package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}
	t1 := t.In(location) // 转换到纽约时区
	fmt.Println(t1)
	fmt.Println(t == t1)
}
