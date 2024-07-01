package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 `case` 语句中，逗号来分隔多个表达式。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("休息日")
	default:
		fmt.Println("工作日")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种方式。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before")
	default:
		fmt.Println("after")
	}

	// 类型开关 (`type switch`) 比较类型而非值。可以用来发现一个接口值的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("Don't konw type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
