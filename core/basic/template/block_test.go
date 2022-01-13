package template

import (
	"go-learn/util"
	"log"
	"strings"
	"testing"
	"text/template"
)

// 参考自：text/template/example_test.go
// 语法体现：block + "" + range + println + define（Parse twice）+ join（Inject Func）
const (
	master = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	// join 不是内置的，需要在构建 template 时指定
	overlay = `{{define "list"}} {{join . ", "}}{{end}}`
)

func TestComplex(t *testing.T) {
	funcs := template.FuncMap{"join": strings.Join}
	data := []string{"A", "B", "C"}

	// 为即将 Parse 要用的模板准备外置函数
	// Parse：完美体现了 block 的语义，为这个 block 引入模板，则该 block 就不会定义模板而是直接使用该模板了
	handleTmpl, err := template.New("test").Funcs(funcs).Parse(master)
	if err != nil {
		t.Fatal(err)
	}
	result, err := util.AsString(handleTmpl, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)

	// three
	overlayTmpl, err := template.Must(handleTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	result, err = util.AsString(overlayTmpl, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
