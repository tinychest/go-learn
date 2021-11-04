package model

import "github.com/xuri/excelize/v2"

type Sheet struct {
	Name string
	Rows []IRow
}

func (s *Sheet) Fill(file *excelize.File) {
	var index = 1
	for _, row := range s.Rows {
		index = row.Fill(file, s.Name, index)
	}
}
