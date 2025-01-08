package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	//打开XML文件
	file, err := os.Open("./info.xml")
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return
	}
	defer file.Close()
	info := website{}
	//创建XML解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
