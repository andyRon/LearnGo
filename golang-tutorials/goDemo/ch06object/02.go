package main

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// 类的初始化函数
//
//	func NewStudent(id uint, name string, male bool, score float64) *Student {
//		return &Student{id: id, name: name, male: male, score: score}
//	}
func NewStudent(id uint, name string, score float64) *Student {
	return &Student{id: id, name: name, score: score}
}

// 定义成员方法
func (s Student) GetName() string {
	return s.name
}
func (s *Student) SetName(name string) {
	s.name = name
}

// 修改默认的Go版toString
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}", s.id, s.name, s.male, s.score)
}

func main() {
	student := NewStudent(1, "andyron", 100)
	fmt.Println(student)
	fmt.Println(student.GetName())
	student.SetName("小狗")
	fmt.Println(student.GetName())
}
