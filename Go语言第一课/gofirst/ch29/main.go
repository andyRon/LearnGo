package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err error = 1

	var err error
	err = errors.New("error1")
	fmt.Printf("%T\n", err)
}
