package main

import (
	"compress/gzip"
	"log"
	"os"
)

// 压缩

func main() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	// 当写入gizp writer数据时,它会依次压缩数据并写入底层的文件中
	// 不必关心它是如何压缩的,像普通的writer一样操作即可
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("已经压缩数据并写入文件.")
}
