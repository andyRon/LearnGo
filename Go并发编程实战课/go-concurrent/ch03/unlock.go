package main

import (
	"fmt"
	"sync"
)

func main() {
	foo()
}

func foo() {
	var mu sync.Mutex
	defer mu.Unlock() // panic: sync: unlock of unlocked mutex
	fmt.Println("hello world")
}
