package template

import (
	"go-learn/tool"
	"testing"
)

const tWith = `
原来的点：{{.}}
{{- with .Word}}
现在的点：{{.}}
{{end}}`

func TestWith(t *testing.T) {
	data := map[string]interface{}{"Word": "WITH"}

	t.Log(tool.MustRenderString(tWith, data))
}
