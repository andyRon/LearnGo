package main

// 自定义泛型类型

type Set[T comparable] map[T]struct{}

type sliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

type Map[K, V any] struct {
	//root *node[K, V]
	compare func(K, K) int
}

type element[T any] struct {
	next *element[T]
	val  T
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type NumericAbs[T Numeric] interface {
	Abs() T
}
