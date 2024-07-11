package main

func main() {
	defer Trace("main")()
	foo()
}

func Trace(name string) func() {
	println("enter: ", name)
	return func() {
		println("exit: ", name)
	}
}

func foo() {
	defer Trace("foo")()
	bar()
}

func bar() {
	defer Trace("bar")()
}

/*
enter:  main
enter:  foo
enter:  bar
exit:  bar
exit:  foo
exit:  main

*/
