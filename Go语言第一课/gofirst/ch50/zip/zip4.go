package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

// 解压
func main() {
	// 打开一个gzip文件
	// 文件是一个reader,但是可以使用各种数据源,比如Web服务器返回的gzipped内容,它的内容不是一个文件,而是一 个内存流
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()
	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()
	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
