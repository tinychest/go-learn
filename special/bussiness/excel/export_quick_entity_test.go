package excel

import (
	"testing"
)

type Fruit struct {
	Name  string `title:"名称"`
	Color string `title:"颜色"`
	Size  string `title:"大小"`
}

func TestQuickEntityWriteToFile(t *testing.T) {
	body := []interface{}{
		&Fruit{"香蕉", "黄色", "10"},
		&Fruit{"草莓", "红色", "4"},
		&Fruit{"西瓜", "绿色", "50"},
	}
	err := QuickWriteToFileWithEntity("quick_entity", body)
	if err != nil {
		panic(err)
	}
}
