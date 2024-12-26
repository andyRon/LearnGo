package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

// Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

// Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

// Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

/*
	interface可以被任意的对象实现。我们看到上面的Men interface被Human、Student和Employee实现。
	同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。

	任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。
*/
