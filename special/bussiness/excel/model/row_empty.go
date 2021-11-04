package model

import (
	"github.com/xuri/excelize/v2"
)

type RowEmpty struct {
	Height float64
}

func (r *RowEmpty) Fill(file *excelize.File, sheetName string, rowIndex int) int {
	r.SetHeight(file, sheetName, r.Height, rowIndex)
	return rowIndex + 1
}

func (r *RowEmpty) SetHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64 {
	curHeight, _ := file.GetRowHeight(sheetName, rowIndex)
	if height == 0 || height <= curHeight {
		height = curHeight
		return curHeight
	}
	if err := file.SetRowHeight(sheetName, rowIndex, height); err != nil {
		panic(err)
	}
	r.Height = height
	return height
}

func (r *RowEmpty) FromData() IRow {
	return &RowEmpty{}
}