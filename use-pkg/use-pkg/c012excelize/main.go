package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	//w()
	//r()
	//u()

	ys()
}

func w() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", 100)
	f.SaveAs("c012excelize/Book1.xlsx") // 保存文件
}

func r() {
	f, _ := excelize.OpenFile("c012excelize/Book1.xlsx")
	defer f.Close() // 必须关闭！

	cell, _ := f.GetCellValue("Sheet1", "A1") // 获取单元格值
	rows, _ := f.GetRows("Sheet1")            // 获取所有行数据

	fmt.Println(cell, rows)
}

func u() {
	f, _ := excelize.OpenFile("c012excelize/Book1.xlsx")
	defer f.Close()

	f.SetCellValue("Sheet1", "A1", "Updated!") // 修改单元格
	f.DeleteSheet("Sheet2")                    // 删除工作表
	f.SetSheetName("Sheet1", "Data")           // 重命名表
	f.SaveAs("c012excelize/updated.xlsx")      // 另存为新文件
}

// 样式设置
func ys() {
	f, _ := excelize.OpenFile("c012excelize/Book1.xlsx")
	defer f.Close()

	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FF0000"},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.SetCellStyle("Sheet1", "A1", "A1", style) // 应用样式
}
