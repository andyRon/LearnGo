package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "persimmon", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "andy", age: 18},
		{name: "tom", age: 17},
		{name: "lu", age: 23},
	}
	slices.SortFunc(people, func(a, b Person) int {
		//return cmp.Compare(a.age, b.age)
		return cmp.Compare(len(a.name), len(b.name))
	})
	fmt.Println(people)
}
