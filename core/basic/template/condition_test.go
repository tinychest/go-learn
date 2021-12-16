package template

import (
	"fmt"
	"go-learn/util"
	"testing"
)

// and 和 or 分别对应了 template.and 和 template.and

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

	fmt.Println(util.MustRenderString(tIf, data))
}

// if A then B else A
const tAnd = `{{and .A .B}}`
func TestCondition2(t *testing.T) {
	data := map[string]string{
		"A": "1",
		"B": "2",
	}

	fmt.Println(util.MustRenderString(tAnd, data))
}