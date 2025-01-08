package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	// 实例化对象
	info := Website{"Go语言", "https://pkg.go.dev/encoding/xml", []string{"Go语言文档", "Go语言-XML"}}
	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer f.Close()
	// 序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}
