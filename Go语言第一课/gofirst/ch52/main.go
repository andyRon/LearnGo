package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	fmt.Println(get_md5("1"))
	fmt.Println(get_sha1("1"))
	fmt.Println(get_sha256("1"))
}

// md5加密后长度为128位，16字节，返回32位十六进制字符串
func get_md5(str string) string {
	// 创建一个MD5哈希对象
	hash := md5.New()
	// 将字符串写入哈希对象
	io.WriteString(hash, str)
	// 计算MD5哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return fmt.Sprintf("%x", hashBytes)
}

// sha1加密后长度为160位，20字节，返回40位十六进制字符串
func get_sha1(str string) string {
	// 创建一个SHA1哈希对象
	hash := sha1.New()
	// 将字符串写入哈希对象
	io.WriteString(hash, str)
	// 计算SHA1哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return fmt.Sprintf("%x", hashBytes)
}

// sha256加密后长度为256位，32字节，返回64位十六进制字符串
func get_sha256(str string) string {
	// 创建一个SHA256哈希对象
	hash := sha256.New()
	// 将字符串写入哈希对象
	io.WriteString(hash, str)
	// 计算SHA256哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	return fmt.Sprintf("%x", hashBytes)
}
