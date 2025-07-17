package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// 通过一个简单的构建SQL查询语句的例子来直观感受Go反射的“魔法”

func main() {
	stmt, err := ConstructQueryStmt(&Product{})
	if err != nil {
		fmt.Println("construct query stmt for Product error: ", err)
		return
	}
	fmt.Println(stmt)

	stmt, err = ConstructQueryStmt(Person{})
	if err != nil {
		fmt.Println("construct query stmt for Person error: ", err)
		return
	}
	fmt.Println(stmt)
}

// ConstructQueryStmt 参数是结构体实例，得到的是该结构体对应的表的数据查询语句文本
// 采用了一种ORM（Object Relational Mapping，对象关系映射）风格的实现。
// ConstructQueryStmt通过反射获得传入的参数obj的类型信息，包括（导出）字段数量、字段名、字段标签值等，并根据这些类型信息生成SQL查询语句文本。如果结构体字段带有orm标签，该函数会使用标签值替代默认列名（字段名）。
func ConstructQueryStmt(obj any) (stmt string, err error) {
	// 仅支持struct或struct指针类型
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		err = errors.New("only struct is supported")
		return
	}

	buffer := bytes.NewBufferString("")
	buffer.WriteString("SELECT ")

	if typ.NumField() == 0 {
		fmt.Errorf("the type[%s] has no fields", typ.Name())
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if i != 0 {
			buffer.WriteString(", ")
		}
		column := field.Name
		if tag := field.Tag.Get("orm"); tag != "" {
			column = tag
		}

		buffer.WriteString(column)
	}

	stmt = fmt.Sprintf("%s FROM %s", buffer.String(), typ.Name())
	return
}

type Product struct {
	ID        uint32
	Name      string
	Price     uint32
	LeftCount uint32 `orm:"left_count"`
	Batch     string `orm:"batch_number"`
	Updated   time.Time
}

type Person struct {
	ID      string
	Name    string
	Age     uint32
	Gender  string
	Addr    string `orm:"address"`
	Updated time.Time
}
