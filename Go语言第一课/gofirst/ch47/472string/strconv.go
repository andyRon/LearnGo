package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Sprintf("%d", i)
	}
	fmt.Println(time.Now().Sub(startTime))

	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		strconv.Itoa(i)
	}
	fmt.Println(time.Now().Sub(startTime))
}
