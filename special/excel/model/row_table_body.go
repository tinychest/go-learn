package model

import (
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/style"
	"go-learn/special/bussiness/excel/util"
)

type RowTableBody struct {
	RowStyle  *excelize.Style
	RowHeight float64
	Rows      []*RowTableBodyRow
}

type RowTableBodyRow struct {
	style  *excelize.Style
	height float64
	Cells  []*Cell
}

func (r *RowTableBody) Fill(file *excelize.File, sheetName string, rowIndex int) int {
	if r.RowStyle == nil {
		r.RowStyle = style.DefaultTableBodyRow
	}
	r.SetHeight(file, sheetName, r.RowHeight, rowIndex)

	for _, row := range r.Rows {
		rowIndex = row.Fill(file, sheetName, rowIndex)
	}
	return rowIndex
}

func (r *RowTableBodyRow) Fill(file *excelize.File, sheetName string, rowIndex int) int {
	if r.style == nil {
		r.style = style.DefaultTableBodyRow
	}

	var colIndex int
	for i, cell := range r.Cells {
		colIndex = i + 1

		if cell.Style == nil {
			cell.Style = r.style
		}
		cell.fill(file, sheetName, rowIndex, colIndex)
	}
	return rowIndex + 1
}

func (r *RowTableBody) SetHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64 {
	if height == 0 {
		height = util.AutoHeight(r.RowStyle.Font.Size)
	}
	for _, row := range r.Rows {
		if row.style == nil {
			row.style = r.RowStyle
		}
		height = row.setHeight(file, sheetName, height, rowIndex)
	}
	r.RowHeight = height
	return height
}

func (r *RowTableBodyRow) setHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64 {
	r.height = setRowHeight(file, sheetName, rowIndex, r.style.Font.Size, height)
	return r.height
}

func (r *RowTableBody) FromData(rows ...[]interface{}) IRow {
	bodyRows := make([]*RowTableBodyRow, 0, len(rows))

	for _, row := range rows {
		bodyRows = append(bodyRows, &RowTableBodyRow{Cells: (*Cell).FromData(nil, row...)})
	}
	return &RowTableBody{Rows: bodyRows}
}
