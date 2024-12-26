package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	p := reflect.ValueOf(&x)
	v2 := p.Elem()
	v2.SetFloat(4.5)
	fmt.Println("v2 is:", v2.Interface())
}
