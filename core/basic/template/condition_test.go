package template

import (
	"fmt"
	"go-learn/util"
	"testing"
)

const tIf = `
{{if .IsTrue}}
A
{{- else}}
B
{{- end}}`

func TestCondition(t *testing.T) {
	data := map[string]interface{}{
		"IsTrue": true,
		"Word":   "Monster",
	}

	fmt.Println(util.MustRender(tIf, data))
}
