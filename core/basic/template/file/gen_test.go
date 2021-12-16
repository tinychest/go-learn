package file

import (
	"fmt"
	"go-learn/util"
	"testing"
	"text/template"
)

func TestGen(t *testing.T) {
	// go embed
	data := map[string]interface{}{
		"title": "首页",
	}
	fmt.Println(util.MustRenderString(page, data))
}

func TestGen2(t *testing.T) {
	// 读取文件
	data := map[string]interface{}{
		"title": "首页",
	}
	tmpl, err := template.ParseFiles("page.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	result, err := util.AsString(tmpl, data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
