package main

import (
	"fmt"
	"reflect"
)

func main() {
	//r1()
	//r2()
	//r3()
	//r4()
	//r_elem()
	//r_field()
	r_modify()
}

func r1() {
	var intNum int
	fmt.Println(reflect.TypeOf(intNum))
	typeOfIntNum := reflect.TypeOf(intNum)                // Type
	fmt.Println(typeOfIntNum.Name(), typeOfIntNum.Kind()) // int int  获取类型的名称(Name)和类型的种类(Kind)

	type structType struct{}
	typeOfStructType := reflect.TypeOf(structType{})
	fmt.Println(typeOfStructType.Name(), typeOfStructType.Kind()) // structType struct
}

func r2() {
	var intNum int = 100
	fmt.Println(reflect.ValueOf(intNum))
}

func r3() {
	var intNum int = 100
	valueOfintNum := reflect.ValueOf(intNum)
	// 将类型强制转换为int
	var originIntNum int64 = int64(valueOfintNum.Int())
	fmt.Println(originIntNum)
}

func r4() {
	var intNums = []int{1, 3, 4, 7, 9}
	valueOfIntNums := reflect.ValueOf(intNums)
	fmt.Println(valueOfIntNums.IsNil())
	fmt.Println(valueOfIntNums.IsValid())
}

func r_elem() {
	intNum := 200
	ins := &intNum
	// 获取指针变量的类型和种类
	typeOfIntNumPtr := reflect.TypeOf(ins)
	fmt.Println(typeOfIntNumPtr.Name(), typeOfIntNumPtr.Kind())
	// 获取指针变量所表示的变量的类型名和种类
	typeOfIntNumPtr = typeOfIntNumPtr.Elem()
	fmt.Println(typeOfIntNumPtr.Name(), typeOfIntNumPtr.Kind())
}

func r_field() {
	type cat struct {
		name string `meaning:"全名"`
		age  int    `meaning:"年龄"`
	}
	catOne := cat{name: "小戎", age: 18}
	typeOfCat := reflect.TypeOf(catOne)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	valueOfCat := reflect.ValueOf(catOne)
	// 通过反射获取结构体成员的名称和值
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Println(fieldType.Index, fieldType.Name, valueOfCat.Field(i), fieldType.Tag)
	}
	// 查找名称为age的成员
	catType, ok := typeOfCat.FieldByName("age")
	if ok {
		// 输出age成员的部分Tag文本
		fmt.Println(catType.Tag.Get("meaning"))
	}
}

func r_modify() {
	numA := 100
	// 获取numA的地址
	addrValueOfNumA := reflect.ValueOf(&numA)
	fmt.Println(&numA)
	// 获取numA地址指向的值，即numA的值
	valueOfNumA := addrValueOfNumA.Elem()
	// 获取valueOfNumA的地址
	fmt.Println(valueOfNumA.Addr())
	// 判断该值是否可被寻址
	fmt.Println(valueOfNumA.CanAddr())

}