package main

// go build -gcflags="-m" escape.go

func main() {
	var m = make([]int, 10240)
	println(m[0])
}
