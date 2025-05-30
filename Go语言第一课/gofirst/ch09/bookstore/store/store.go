package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct {
	// 每个字段都使用了json标签，指定了在JSON序列化和反序列化时对应的键名。
	Id      string   `json:"id"`      // 图书ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
