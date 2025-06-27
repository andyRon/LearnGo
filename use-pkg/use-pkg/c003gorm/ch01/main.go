package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "D43", Price: 150})

	var product Product
	//db.First(&product, 1)
	db.First(&product, "code = ?", "D43")

	//
	//// 更新
	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//db.Delete(&product, 1)

	fmt.Println(product)

	var products []Product
	result := db.Find(&products)

	fmt.Println(result.RowsAffected)

}
