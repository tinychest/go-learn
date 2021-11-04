package excel

import (
	"errors"
	"go-learn/special/bussiness/excel/model"
	"io"
	"reflect"
)

const (
	defaultSheetName = "Sheet1"
	tableHeadTagName = "title"
)

func QuickExportWithEntity(fileName string, entities []interface{}) (io.WriterTo, error) {
	excel, err := entityToExcel(fileName, entities)
	if err != nil {
		return nil, err
	}
	return excel.Export()
}

func QuickWriteToFileWithEntity(fileName string, entities []interface{}) error {
	excel, err := entityToExcel(fileName, entities)
	if err != nil {
		return err
	}
	return excel.WriteToFile()
}

func QuickExport(fileName string, head []string, body ...[]interface{}) (io.WriterTo, error) {
	return dataToExcel(fileName, head, body...).Export()
}

func QuickWriteToFile(fileName string, head []string, body ...[]interface{}) error {
	return dataToExcel(fileName, head, body...).WriteToFile()
}

func entityToExcel(fileName string, entities []interface{}) (model.IExcel, error) {
	head, body ,err := entityToData(entities)
	if err != nil {
		return nil, err
	}
	return dataToExcel(fileName, head, body...), nil
}

func entityToData(entities []interface{}) ([]string, [][]interface{}, error) {
	// 从结构体中读取 entities 的 <标签> 作为表格头，没有标签的不作为表格导出数据
	if len(entities) < 1 {
		return nil, nil, errors.New("entities can not be empty")
	}

	// HEAD
	entity := entities[0]
	v := reflect.ValueOf(entity).Elem()
	t := reflect.TypeOf(entity).Elem()

	var (
		tagIndex = make([]int, 0, t.NumField())
		tagValue = make([]string, 0, t.NumField())
	)
	for i := 0; i < t.NumField(); i++ {
		headValue := t.Field(i).Tag.Get(tableHeadTagName)
		if len(headValue) == 0 {
			break
		}
		tagIndex = append(tagIndex, i)
		tagValue = append(tagValue, headValue)
	}
	if len(tagValue) == 0 {
		return nil, nil, errors.New("no valid struct field (forgotten add tag?)")
	}

	// BODY
	var (
		rowLen = len(entities)
		colLen = len(tagIndex)
		body   = make([][]interface{}, 0, rowLen)
	)
	for _, entity = range entities {
		v = reflect.ValueOf(entity).Elem()

		fieldValues := make([]interface{}, 0, colLen)
		for _, index := range tagIndex {
			fieldValues = append(fieldValues, v.Field(index).Interface())
		}
		body = append(body, fieldValues)
	}

	return tagValue, body, nil
}

func dataToExcel(fileName string, head []string, body ...[]interface{}) model.IExcel {
	sheet := &model.Sheet{
		Name: defaultSheetName,
		Rows: []model.IRow{
			(*model.RowTableHead).FromData(nil, head...),
			(*model.RowTableBody).FromData(nil, body...),
		},
	}
	return &model.DataExcel{
		Name:        fileName,
		Sheets:      []*model.Sheet{sheet},
		ActiveSheet: 0,
	}
}
