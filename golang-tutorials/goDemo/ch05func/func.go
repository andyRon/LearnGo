package main

import (
	"errors"
	"fmt"
	"reflect"
)

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value.")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func add(a, b *int) (int, error) {
	if *a < 0 || *b < 0 {
		err := errors.New("只支持非负整数相加")
		return 0, err
	}
	*a *= 2
	*b *= 3
	return *a + *b, nil
}

func main() {
	//myPrintf(1, "1", [1]int{1}, true)

	x, y := -1, 2
	z, err := add(&x, &y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
