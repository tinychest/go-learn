package model

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-learn/util/set"
	"io"
	"net/http"
	"net/url"
	"os"
)

type DataExcel struct {
	Name        string
	Base        *excelize.File
	Sheets      []*Sheet
	ActiveSheet int
}

func (de *DataExcel) valid() error {
	// check excel name empty
	if len(de.Name) == 0 {
		return errors.New("excel name can not be empty")
	}

	// check export sheets
	if len(de.Sheets) == 0 {
		return errors.New("excel sheets can not be empty")
	}

	// check activeSheet
	if de.ActiveSheet < 0 {
		return errors.New("excel activeActive can not be negative")
	}

	// check sheet name dup、empty
	var l = len(de.Sheets)
	if de.Base != nil {
		l += len(de.Base.GetSheetList())
	}
	s := set.NewString(l)
	if de.Base != nil {
		s.Add(de.Base.GetSheetList()...)
	}
	for i, sheet := range de.Sheets {
		if len(sheet.Name) == 0 {
			return fmt.Errorf("sheet name empty, index:%d", i)
		}
		s.Add(sheet.Name)
	}
	if s.Len() != l {
		return errors.New("exists duplicate sheet name")
	}
	return nil
}

// Export 导出
func (de *DataExcel) Export() (io.WriterTo, error) {
	if err := de.valid(); err != nil {
		return nil, err
	}

	if de.Base == nil {
		de.Base = excelize.NewFile()
		// 默认会创建一个 sheet；NewSheet 是幂等的
		de.Base.SetSheetName(de.Base.GetSheetName(0), de.Sheets[0].Name)
	}
	// 填充数据
	for _, sheet := range de.Sheets {
		de.Base.NewSheet(sheet.Name)
		sheet.Fill(de.Base)
	}
	// 设置默认展示的 Sheet
	de.Base.SetActiveSheet(de.ActiveSheet)

	return de.Base, nil
}

func (de *DataExcel) Response(headSetter func(string, string), writer http.ResponseWriter) error {
	writeTo ,err := de.Export()
	if err != nil {
		return err
	}

	headSetter("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	headSetter("Content-Disposition", "attachment; filename="+url.PathEscape(de.Name))
	_, err = writeTo.WriteTo(writer)
	return err
}

func (de *DataExcel) WriteToFile() error {
	writeTo ,err := de.Export()
	if err != nil {
		return err
	}

	if len(de.Name) == 0 {
		return errors.New("empty file name")
	}

	f, err := os.Create(de.Name + ".xlsx")
	if err != nil {
		return err
	}
	_, err = writeTo.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}
