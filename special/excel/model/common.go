package model

import (
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/util"
)

func setRowHeight(file *excelize.File, sheetName string, rowIndex int, fontSize, height float64) float64 {
	curHeight, _ := file.GetRowHeight(sheetName, rowIndex)
	if height == 0 {
		height = util.AutoHeight(fontSize)
	}
	if height <= curHeight {
		height = curHeight
		return curHeight
	}
	if err := file.SetRowHeight(sheetName, rowIndex, height); err != nil {
		panic(err)
	}
	return height
}
