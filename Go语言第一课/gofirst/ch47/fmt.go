package main

import "fmt"

func main() {
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)

	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)

	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	s := "枯藤"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)

	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}
