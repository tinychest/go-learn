package model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/style"
)

type RowTableHead struct {
	*excelize.Style
	Height   float64
	Cells    []*Cell
	NotFloat bool
}

func (r *RowTableHead) Fill(file *excelize.File, sheetName string, rowIndex int) int {
	if r.Style == nil {
		r.Style = style.DefaultTableHead
	}
	r.SetHeight(file, sheetName, r.Height, rowIndex)
	if !r.NotFloat {
		err := file.SetPanes(sheetName, fmt.Sprintf(`{"freeze": true, "y_split": %d, "top_left_cell": "A%d"}`, rowIndex, rowIndex + 1))
		if err != nil {
			panic(err)
		}
	}

	var colIndex int
	for i, cell := range r.Cells {
		colIndex = i + 1

		if cell.Style == nil {
			cell.Style = r.Style
		}
		cell.fill(file, sheetName, rowIndex, colIndex)
	}
	return rowIndex + 1
}

func (r *RowTableHead) SetHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64 {
	r.Height = setRowHeight(file, sheetName, rowIndex, r.Font.Size, height)
	return r.Height
}

func (r *RowTableHead) FromData(list ...string) IRow {
	return &RowTableHead{Cells: (*Cell).FromHeadData(nil, list...)}
}
