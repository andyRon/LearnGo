package main

import (
	. "c003gorm/util"
	"fmt"
	"time"
)

func main() {
	//create1()
	//create2()
	//create3()
	//create4()
	create5()
}

func create1() {
	user := User{Name: "andyron", Age: 18, Birthday: time.Now()}
	DB.AutoMigrate(&User{}) // 自动迁移就是创建表

	res := DB.Create(&user)

	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
	fmt.Println(user.ID)
}

func create2() {
	users := []*User{
		{Name: "Andy", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	DB.AutoMigrate(&User{})

	res := DB.Create(users)
	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
}

func create3() {
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	DB.Select("Name", "Age").Create(&user)

	res := DB.Create(&user)

	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
	fmt.Println(user.ID)

	//DB.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
}

func create4() {
	var users = []User{{Name: "andy1"}, {Name: "andy2"}, {Name: "andy3"}}
	DB.Create(&users)

	for _, user := range users {
		println(user.ID)
	}
}

func create5() { // TODO
	var users = []User{{Name: "andy11"}, {Name: "andy21"}, {Name: "andy31"}, {Name: "andy41"}, {Name: "andy51"}, {Name: "andy61"}}
	DB.CreateInBatches(users, 5)

	for _, user := range users {
		println(user.ID)
	}
}
