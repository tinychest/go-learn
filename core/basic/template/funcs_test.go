package template

import (
	"go-learn/tool"
	"strings"
	"testing"
	"text/template"
)

// 标准库函数 len
// 自定义注册函数 strings.Join
// 详见 template 包注释
const tF = `len:{{len .}} cap:{{cap .}} {{join . ", "}}`

func TestFuncs(t *testing.T) {
	data := []string{"A", "B", "C"}
	funcs := template.FuncMap{"join": strings.Join}

	tpl, err := template.New("test").Funcs(funcs).Parse(tF)
	if err != nil {
		t.Fatal(err)
	}
	res, err := tool.AsString(tpl, data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
