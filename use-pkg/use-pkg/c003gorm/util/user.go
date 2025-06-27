package util

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func (table *User) TableName() string {
	return "user"
}
