package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

/*
$ go run ch82-exit.go
exit status 3


$ go build ch82-exit.go
$ ./ch82-exit
$ echo $?
3
*/
