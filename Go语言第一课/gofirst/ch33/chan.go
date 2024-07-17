package main

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}

func t1() {
	ch1 := make(chan int)
	ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	n := <-ch1
	println(n)
}

func t2() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13 // 将发送操作放入一个新goroutine中执行
	}()
	n := <-ch1
	println(n)
}

func t3() {
	ch2 := make(chan int, 1)
	n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起
	println(n)

	ch3 := make(chan int, 1)
	ch3 <- 17 // 向ch3发送一个整型数17
	ch3 <- 27 // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
}

func t4() {
	ch1 := make(chan<- int, 1) // 只发送channel类型
	ch2 := make(<-chan int, 1) // 只接收channel类型

	<-ch1     // invalid operation: <-ch1 (receive from send-only type chan<- int)
	ch2 <- 13 // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
}
