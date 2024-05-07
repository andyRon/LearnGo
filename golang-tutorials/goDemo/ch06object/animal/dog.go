package animal

type Dog struct {
	// 相当于继承
	annimal *Animal
	pet     Pet
}

func NewDog(animal *Animal, pet Pet) Dog {
	return Dog{annimal: animal, pet: pet}
}

func (d Dog) Call() string {
	return d.annimal.Call() + "汪汪汪"
}
func (d Dog) FavorFood() string {
	return d.annimal.FavorFood() + "骨头和鱼"
}

func (d Dog) GetName() string {
	return d.pet.GetName()
}
