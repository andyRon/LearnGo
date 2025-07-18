package main

import (
	"github.com/andyron/workerpool1"
	"time"
)

func main() {
	p := workerpool1.New(5)

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			println("task: ", i, "err: ", err)
		}
	}
	p.Free()
}
