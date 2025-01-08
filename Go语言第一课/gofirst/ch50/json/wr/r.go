package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []website
	//创建JSON解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
