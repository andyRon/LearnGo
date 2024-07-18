package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"a", "n", "d", "y"}
	slices.Sort(strs)
	fmt.Println(strs)

	ints := []int{1, 5, 8, 7}
	slices.Sort(ints)
	fmt.Println(ints)

	fmt.Println(slices.IsSorted(ints))
}
