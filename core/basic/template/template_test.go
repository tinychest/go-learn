package template

import (
	"fmt"
	"go-learn/util"
	"testing"
)

// 注意，当前是 template 中，模板的概念

// 要求引用的模板一定存在，否则：template "y" not defined

const tT = `
{{- define "x"}}x{{end}}
{{- template "y"}}`

func TestTemplate(t *testing.T) {
	var data interface{}

	fmt.Println(util.MustRender(tT, data))
}
