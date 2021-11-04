package model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/util"
)

type Cell struct {
	*excelize.Style
	Width float64
	Data  interface{}
}

func (c *Cell) fill(file *excelize.File, sheetName string, rowIndex, colIndex int) {
	coordinate, _ := excelize.CoordinatesToCellName(colIndex, rowIndex)

	// 样式
	styleId, _ := file.NewStyle(c.Style)
	if err := file.SetCellStyle(sheetName, coordinate, coordinate, styleId); err != nil {
		panic(err)
	}
	c.setWidth(file, sheetName, c.Width, colIndex)

	// 数据
	if err := file.SetCellValue(sheetName, coordinate, c.Data); err != nil {
		panic(err)
	}
}

func (c *Cell) setWidth(file *excelize.File, sheetName string, width float64, colIndex int) {
	c.Width = width

	colName, _ := excelize.ColumnNumberToName(colIndex)
	curWidth, _ := file.GetColWidth(sheetName, colName)
	if c.Width == 0 {
		c.Width = util.AutoWidth(c.Font.Size, fmt.Sprintf("%s", c.Data))
	}
	if c.Width <= curWidth {
		c.Width = curWidth
		return
	}
	if err := file.SetColWidth(sheetName, colName, colName, c.Width); err != nil {
		panic(err)
	}
}

func (c *Cell) FromHeadData(list ...string) []*Cell {
	cells := make([]*Cell, 0, len(list))
	for _, item := range list {
		cells = append(cells, &Cell{Data: item})
	}
	return cells
}

func (c *Cell) FromData(list ...interface{}) []*Cell {
	cells := make([]*Cell, 0, len(list))
	for _, item := range list {
		cells = append(cells, &Cell{Data: item})
	}
	return cells
}