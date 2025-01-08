package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 写JSON文件

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	info := []Website{{"JSON", "https://go.dev/", []string{"https://pkg.go.dev/encoding/json", "https://pkg.go.dev/encoding"}},
		{"XML", "https://pkg.go.dev/encoding/xml", []string{"https://pkg.go.dev/encoding/xml", "https://go.dev/doc/effective_go"}}}
	// 创建文件
	filePtr, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建JSON编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}
