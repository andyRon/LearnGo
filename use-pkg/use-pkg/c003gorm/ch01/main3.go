package main

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      uint
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

type Blog2 struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"` // 为 db 中的字段名添加前缀
	Upvotes int32
}
