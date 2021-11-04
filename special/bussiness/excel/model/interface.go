package model

import (
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
)

type ICell interface {
	fill(file *excelize.File, sheetName string, rowIndex, colIndex int)
	setWidth(file *excelize.File, sheetName string, width float64, colIndex int)
}

type IRow interface {
	Fill(file *excelize.File, sheetName string, index int) int
	SetHeight(file *excelize.File, sheetName string, height float64, rowIndex int) float64
}

type ISheet interface {
	Fill(file *excelize.File)
}

type IExcel interface {
	Export() (io.WriterTo, error)
	Response(headSetter func(string, string), w http.ResponseWriter) error
	WriteToFile() error
}