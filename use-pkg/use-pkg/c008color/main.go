package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
)

func main() {

	// 1️⃣基础样式
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	color.White("This is white")
	color.HiWhite("This is hiwhite") // 函数名前加Hi前缀，为加深颜色

	fmt.Println("original")
	b := color.New(color.Bold)
	b.Println("bold")

	// 2️⃣其他样式
	color.New(color.Italic).Println("样式")
	color.New(color.Underline).Println("样式")
	color.New(color.CrossedOut).Println("删除线样式")
	color.New(color.Faint).Println("模糊样式")

	// 3️⃣组合
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline) // 景色Fg、背景色Bg
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	// 白色背景，红色文字
	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	fmt.Println(color.FgHiBlue)

	// 4️⃣原有打印的改造
	// Use handy standard colors
	color.Set(color.FgYellow)
	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")
	color.Unset() // Don't forget to unset
	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset() // Use it in your function

	fmt.Println("All text will now be bold magenta.")

	// 5️⃣输出接口
	myWriter := io.Writer(color.Output)
	color.New(color.FgBlue).Fprintln(myWriter, "blue color!")
	blue := color.New(color.FgBlue)
	blue.Fprint(myWriter, "This will print text in blue.")

	// 6️⃣自定义函数名
	fmt.Println()

	myred := color.New(color.FgRed).PrintfFunc()
	myred("Warning")

	// 7️⃣禁用颜色  TODO
	//flag.Parse() // 需要先解析flag
	//// 运行时添加`--no-color`参数即可
	//var flagNoColor = flag.Bool("no-color", false, "Disable color output")
	//if *flagNoColor {
	//	color.NoColor = true   // disables colorized output
	//	fmt.Println("颜色输出已禁用") // 改为普通打印
	//
	//}
	// 对单一颜色对象的禁用
	c = color.New(color.FgCyan)
	c.Println("Prints cyan text")
	c.DisableColor()
	c.Println("This is printed without any color")
	c.EnableColor()
	c.Println("This prints again cyan...")

	fmt.Println()

	// 8️⃣命令行参数  go run main.go -color=red
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
