package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//fmt.Println(os.ModeDir)
	//fmt.Println(reflect.TypeOf(os.ModeDir).Name())
	//f()
	//f1()
	//f2()
	//f3()
	//f5_1()
	//f5_2()
	//f5_3()
	//f6()
	//f6_1()
	f7()
}

func f() {
	os.Mkdir("goDir", 0777)
	os.MkdirAll("goDir/dir1/dir2", 0777)
	err := os.Remove("goDir")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("goDir")
}

var (
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "test/test2/"
	fileName = "demo.txt"
	filePath = path + fileName
)

func f1() {
	err = os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("成功创建目录")
	}
	// 创建空白文件
	newFile, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newFile)
	newFile.Close()
	// 查看文件信息
	fileInfo, err = os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		log.Fatal("文件不存在")
	}
	fmt.Println("文件名称: ", fileInfo.Name())
	fmt.Println("文件大小: ", fileInfo.Size())
	fmt.Println("文件权限:", fileInfo.Mode())
	fmt.Println("最后修改时间: ", fileInfo.ModTime())
	fmt.Println("是否是文件夹: ", fileInfo.IsDir())
	fmt.Printf("系统接口类型: %T\n", fileInfo.Sys())
	fmt.Printf("系统信息: %+v\n\n", fileInfo.Sys())
}

func f2() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func f3() {
	//简单地以只读的方式打开
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	//打印文件内容
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
	// OpenFile提供更多的选项,第一个参数是文件路径,第二个参数是打开时的属性,第三个参数是打开时的文件权限模式
	file, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func f5_1() {
	// 打开原始文件
	originalFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create(filePath + "_copy")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// 从源文件中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("复制了 %d 个字节", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func f5_2() {
	file, _ := os.Open(filePath)
	defer file.Close()
	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5

	// whence的意思，0 文件开始位置，1 当前位置，2 文件结尾处
	whence := 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("移动到位置5：", newPosition)
	// 从当前位置回退两字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("从当前位置回退两字节：", newPosition)
	// 使用下面的方式得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("当前位置：", currentPosition)
	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("转到文件开始位置(0,0)：", newPosition)
}

func f5_3() {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write([]byte("写入字节。\r\n"))
	file.WriteString("写入字符串。\r\n")

	// 打印内容
	file, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
}

func f6() {
	// 测试写权限
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Fatal("错误：没有写入权限")
	} else if os.IsNotExist(err) {
		log.Fatal("错误：文件不存在")
	} else {
		log.Fatal(err)
	}
	file.Close()

	// 测试读权限
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Fatal("错误：没有读取权限")
	} else if os.IsNotExist(err) {
		log.Fatal("错误：文件不存在")
	} else {
		log.Fatal(err)
	}
	file.Close()

}

func f6_1() {
	// 使用Linux风格改变文件权限
	err := os.Chmod(filePath, 0777)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(os.Getuid(), os.Getgid())
	// 改变文件所有者
	err = os.Chown(filePath, os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}
	// 查看文件信息
	fileInfo, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("文件不存在.")
		}
		log.Fatal(err)
	}
	fmt.Println("最后修改时间: ", fileInfo.ModTime())
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes(filePath, lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

func f7() {
	// 创建一个硬链接
	// 创建后同一个文件内容会有两个文件名,改变一个文件的内容会影响另一个
	// 删除和重命名不会影响另一个
	hardLink := filePath + "_hl"
	err := os.Link(filePath, hardLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建硬链接")
	// 创建一个软链接
	softLink := filePath + "_sl"
	err = os.Symlink(fileName, softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建软链接")

	// Lstat返回一个文件的信息,但是当文件是一个软链接时,它返回软链接的信息,而不是引用的文件的信息
	// Symlink在Windows中不工作
	fileInfo, err := os.Lstat(softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("链接信息: %+v", fileInfo)
	// 改变软链接的拥有者不会影响原始文件
	err = os.Lchown(softLink, os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
