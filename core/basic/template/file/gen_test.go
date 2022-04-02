package file

import (
	"go-learn/util"
	"testing"
	"text/template"
)

func TestGen(t *testing.T) {
	// go embed
	data := map[string]interface{}{
		"title": "首页",
	}
	t.Log(util.MustRenderString(page, data))
}

func TestGen2(t *testing.T) {
	// 读取文件
	data := map[string]interface{}{
		"title": "首页",
	}
	tpl, err := template.ParseFiles("page.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	res, err := util.AsString(tpl, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
