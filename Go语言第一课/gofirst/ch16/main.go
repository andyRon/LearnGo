package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	m1[1] = "a"
	m1[2] = "b"
	m1[3] = "c"
	m1[4] = "d"
	for k, v := range m1 {
		println(k, v)
	}
	println(len(m1))

	v, ok := m1[6]
	if !ok {
		println("6不在map中")
	} else {
		println(v)
	}

	for k, v := range m1 {
		fmt.Printf("[%d, %s]\n", k, v)
	}
}
