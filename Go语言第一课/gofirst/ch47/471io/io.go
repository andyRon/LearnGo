package main

import (
	"fmt"
	"gofirst/ch47/util"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//ReaderExample()
	//println(util.GetProjectRoot())
	fmt.Println(os.Args)
	fmt.Println(filepath.Dir(os.Args[0]))
}

/*
1 io为 IO 原语（I/O primitives）提供基本的接口
*/
func ReaderExample() {
FOREND:
	for {
		readerMenu()

		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err  error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，以回车结束：")
			data, err = ReadFrom(os.Stdin, 11)
		case "2":
			file, err := os.Open(util.GetProjectRoot() + "01.txt") // TODO
			if err != nil {
				fmt.Println("打开文01.txt件失败：", err)
				continue
			}
			data, err = ReadFrom(file, 9)
			file.Close()
		case "3":
			data, err = ReadFrom(strings.NewReader("from string"), 12)
		case "4":
			fmt.Println("暂未实现！")
		case "b":
			fmt.Println("返回上级菜单！")
			break FOREND
		case "q":
			fmt.Println("退出程序！")
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入！")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败，可以试试从其他输入源读取！")
		} else {
			fmt.Printf("读取到的数据为：%s\n", data)
		}
	}

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readerMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}
