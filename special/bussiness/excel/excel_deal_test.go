package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"testing"
)

// 第三方的 excel 处理类库 - 360 的 excelize

// 还有写入图片的功能，这里就不做拓展了
func TestExcelDeal(t *testing.T) {
	// writeBasic("Test1")
	// read("Test1")

	writeChart("Test2")
}

// 表格是什么，打开生成的 excel 就知道了
func writeChart(filename string) {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}

	f := excelize.NewFile()
	for k, v := range categories {
		_ = f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		_ = f.SetCellValue("Sheet1", k, v)
	}

	// 到这里才是真正的写入表格数据
	if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		println(err)
		return
	}

	if err := f.SaveAs(filename + ".xlsx"); err != nil {
		println(err)
	}
}

func read(filename string) {
	f, _ := excelize.OpenFile(filename + ".xlsx")

	// Get value from cell by given worksheet name and axis.
	cell, _ := f.GetCellValue("Sheet1", "B2")
	println(cell)

	// Get all the rows in the given worksheet.
	rows, _ := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		println()
	}
}

func writeBasic(filename string) {
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
	if err := f.SaveAs(filename + ".xlsx"); err != nil {
		println(err)
	}
}
