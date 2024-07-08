package main

import (
	"fmt"
	"image"
	_ "image/gif"  // 以空导入方式注入gif图片格式驱动
	_ "image/jpeg" // 以空导入方式注入jpeg图片格式驱动
	_ "image/png"  // 以空导入方式注入png图片格式驱动
	"os"
)

func main() {
	width, height, err := imageSize(os.Args[1])
	if err != nil {
		fmt.Println("获取图片大小错误：", err)
		return
	}
	fmt.Printf("图片大小：[%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
	f, _ := os.Open(imageFile) // 打开图文文件
	defer f.Close()

	img, _, err := image.Decode(f) // 对文件进行解码，得到图片实例
	if err != nil {
		return 0, 0, err
	}

	b := img.Bounds() // 返回图片区域
	return b.Max.X, b.Max.Y, nil
}
