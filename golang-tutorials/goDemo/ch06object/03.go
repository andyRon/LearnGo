package main

import (
	"fmt"
	. "goDemo/ch06object/animal" // 使用了 . 作为前缀，无需使用包名前缀 animal. 引用
)

func main() {
	animal := NewAnimal("草狗")
	pet := NewPet("宠物狗")
	dog := NewDog(&animal, pet)

	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())

}
