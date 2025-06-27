package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ProblemBasic struct {
	gorm.Model
	ID         uint   `gorm:"primarykey;" json:"id"`
	Identity   string `gorm:"column:identity;type:varchar(255);" json:"identity"`  // 问题的唯一标识
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`        // 问题标题
	Content    string `gorm:"column:content;type:text;" json:"content"`            // 正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` // 最大运行时长
	MaxMem     int    `gorm:"column:max_mem;type:int(11);" json:"max_mem"`         // 最大运行内存
	PassNum    int64  `gorm:"column:pass_num;type:int(11);" json:"pass_num"`       // 通过次数
	SubmitNum  int64  `gorm:"column:submit_num;type:int(11);" json:"submit_num"`   // 提交次数
	//DeletedAt  gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (ProblemBasic) TableName() string {
	return "problem_basic" // 指定单数形式的表名
}

func main() {
	dsn := "andy:iop654321@tcp(127.0.0.1:3306)/oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	var products []ProblemBasic
	db.Unscoped().First(&products, 2)
	fmt.Println(products)
}
