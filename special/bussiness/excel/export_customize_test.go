package excel

import (
	"bytes"
	_ "embed"
	"github.com/xuri/excelize/v2"
	"go-learn/special/bussiness/excel/model"
	"testing"
)

//go:embed base.xlsx
var Base []byte

var BaseExcel *excelize.File

func init() {
	var err error

	if BaseExcel, err = excelize.OpenReader(bytes.NewReader(Base)); err != nil {
		panic(err)
	}
}

func TestModule(*testing.T) {
	sheet1 := &model.Sheet{
		Name: "sheet01",
		Rows: []model.IRow{
			&model.RowAWord{Cell: &model.Cell{Data: "123"}},
			&model.RowEmpty{},
			&model.RowAWord{Cell: &model.Cell{Data: "456"}},
			&model.RowTableHead{
				Cells: []*model.Cell{
					{Data: "标题1"},
					{Data: "标题2"},
					{Data: "标题3"},
				},
			},
			&model.RowTableBody{
				Rows: []*model.RowTableBodyRow{
					{
						Cells: []*model.Cell{
							{Data: "小明sdfsdfsdfsdfsdf1"},
							{Data: "小明2"},
							{Data: "小明3"},
						},
					},
					{
						Cells: []*model.Cell{
							{Data: "大明明1"},
							{Data: "按时打卡三方接口1111dsdfsssssssssdfsdf111111"},
							{Data: "阿瑟东评价哦"},
						},
					},
				},
			},
		},
	}
	sheet2 := &model.Sheet{
		Name: "sheet02",
		Rows: []model.IRow{
			&model.RowAWord{Cell: &model.Cell{Data: "123"}},
			&model.RowEmpty{},
			&model.RowAWord{Cell: &model.Cell{Data: "456"}},
		},
	}
	excel := model.DataExcel{
		Name:        "test",
		Base:        BaseExcel,
		Sheets:      []*model.Sheet{sheet1, sheet2},
		ActiveSheet: 1,
	}

	if err := excel.WriteToFile(); err != nil {
		panic(err)
	}
}
