package main

import "fmt"

type ordered interface {
	~int | ~int32 | ~int16 | ~int64 | ~int8 |
		~float32 | ~float64 |
		~string |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func maxGenerics[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return max
}

type myString string

func main() {
	var m int = maxGenerics([]int{1, 2, -4, -6, 7, 0})
	fmt.Println(m)                                                           // 输出：7
	fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"}))   // 输出：77
	fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}))  // 输出：7.07
	fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0}))                     // 输出：7
	fmt.Println(maxGenerics([]myString{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
