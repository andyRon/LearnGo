package main

import (
	"flag"
	"fmt"
)

// 命令行标志是为命令行程序指定选项的常用方法。例如，在 wc -l 中， -l 是一个命令行标志。

func main() {

	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// go run ch72-command-line-flags.go -h

/*
$ go run ch72-command-line-flags.go
word: foo
numb: 42
fork: false
svar: bar
tail: []

$ go run ch72-command-line-flags.go -word=opt -numb=7 -fork -svar=flag a c
word: opt
numb: 7
fork: true
svar: flag
tail: [a c]
*/

/*
$ go run ch72-command-line-flags.go -wat
flag provided but not defined: -wat
Usage of /var/folders/8k/ntbhdf615p34cflx1_qwv38r0000gn/T/go-build680355273/b001/exe/ch72-command-line-flags:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
exit status 2
*/
