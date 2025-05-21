package main

import (
	"fmt"
	"sort"
)

// []interface 排序与查找
func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 按年龄升序排序
	//sort.SliceStable(people, func(i, j int) bool { return people[i].Age > people[j].Age }) // 按年龄降序排序
	fmt.Println("Sort by age:", people)

	fmt.Println("Sorted:", sort.SliceIsSorted(people, func(i, j int) bool { return people[i].Age > people[j].Age }))

	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })           // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}
}
