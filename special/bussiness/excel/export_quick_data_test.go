package excel

import (
	"testing"
)

func TestQuickWriteToFile(t *testing.T) {
	head := []string{
		"名称", "颜色", "大小",
	}
	body := [][]interface{}{
		{"香蕉", "黄色", "10"},
		{"草莓", "红色", "4"},
		{"西瓜", "绿色", "50"},
	}
	err := QuickWriteToFile("quick_data", head, body...)
	if err != nil {
		panic(err)
	}
}
