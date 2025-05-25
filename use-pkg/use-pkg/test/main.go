package main

import (
	"flag"
	"github.com/fatih/color"
)

func main() {
	colorArg := flag.String("color", "default", "Text color (blue/red/green/yellow)")
	flag.Parse()
	// 根据参数选择颜色
	switch *colorArg {
	case "blue":
		color.Blue("Hello, Blue World!")
	case "red":
		color.Red("Hello, Red World!")
	case "green":
		color.Green("Hello, Green World!")
	default:
		color.White("Hello, Default World!")
	}
}
