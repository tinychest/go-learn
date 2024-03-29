package template

import (
	"go-learn/tool"
	"testing"
)

// 注意，当前是 template 中，模板的概念

// 要求引用的模板一定存在，否则：template "y" not defined

const tT = `{{/*我是注释*/}}
{{- define "x"}}x{{end}}
{{- template "x"}}`

func TestTemplate(t *testing.T) {
	var data interface{}

	t.Log(tool.MustRenderString(tT, data))
}
