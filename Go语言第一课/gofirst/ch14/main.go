package main

import "fmt"

func main() {
	const (
		Apple, Banana = 11, 22
		Strawberry, Grape
		Pear, Watermelon
	)

	fmt.Println(Strawberry, Grape, Pear, Watermelon)

	const (
		mutexLocked = 1 << iota
		mutexWoken
		mutexStarving
		mutexWaiterShift      = iota
		starvationThresholdNs = 1e6
	)
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)
}
