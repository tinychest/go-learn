package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

// 还有写入图片的功能，这里就不做拓展了
func TestExcelDeal(t *testing.T) {
	writeDemo("test1")
	readDemo("test1")
}

func basic() {
	f := excelize.NewFile()

	// Create a new sheet.
	// 默认会创建一个名为 Sheet1 的表格页
	index := f.NewSheet("Sheet2")

	// Set value of a cell.
	// axis（第二个参数）：[A-Z]正整数 分别代表列标和行标
	_ = f.SetCellValue("Sheet1", "B2", 100)
	_ = f.SetCellValue("Sheet2", "A2", "Hello world.")

	// Set active sheet of the workbook.
	// 设置打开表格，默认展示的表格页
	f.SetActiveSheet(index)

	// 不指定路径，就指定文件名，文件回生成到当前文件所在的目录下
	if err := f.SaveAs("xxx.xlsx"); err != nil {
		println(err)
	}
}

func writeDemo(filename string) {
	rowOne := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large"}
	colOne := map[string]string{"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	body := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	f := excelize.NewFile()
	for k, v := range rowOne {
		_ = f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range colOne {
		_ = f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range body {
		_ = f.SetCellValue("Sheet1", k, v)
	}

	if err := f.SaveAs(filename + ".xlsx"); err != nil {
		println(err)
	}
}

func readDemo(filename string) {
	f, _ := excelize.OpenFile(filename + ".xlsx")

	// Get value from cell by given worksheet name and axis.
	// cell, _ := f.GetCellValue("Sheet1", "B2")
	// println(cell)

	// Get all the rows in the given worksheet.
	rows, _ := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		println()
	}
}
