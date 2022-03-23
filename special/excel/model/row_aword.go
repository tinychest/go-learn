package model

import (
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/style"
)

type RowAWord struct {
	*Cell
	Height float64
}

func (r *RowAWord) Fill(file *excelize.File, sheetName string, rowIndex int) int {
	if r.Style == nil {
		r.Style = style.DefaultRowAWord
	}
	r.SetHeight(file, sheetName, r.Height, rowIndex)

	r.fill(file, sheetName, rowIndex, 1)
	return rowIndex + 1
}

func (r *RowAWord) SetHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64 {
	r.Height = setRowHeight(file, sheetName, rowIndex, r.Font.Size, height)
	return r.Height
}

func (r *RowAWord) FromData(data interface{}) IRow {
	return &RowAWord{
		Cell: &Cell{Data: data},
	}
}
