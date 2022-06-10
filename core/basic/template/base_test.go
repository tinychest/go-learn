package template

import (
	"go-learn/core"
	"go-learn/tool"
	"strings"
	"testing"
	"text/template"
)

const tVar = `Name: {{.Name}}, Age: {{.Age}}`

func TestBase(t *testing.T) {
	params := core.Person{Age: 23, Name: "xiaoming"}

	tpl, err := template.New("test").Parse(tVar)
	if err != nil {
		t.Fatal(err)
	}

	var b = &strings.Builder{}
	if err = tpl.Execute(b, params); err != nil {
		t.Fatal(err)
	}
	t.Log(b.String())
}

func TestVar(t *testing.T) {
	data := core.Person{Age: 23, Name: "xiaoming"}

	t.Log(tool.MustRenderString(tVar, data))
}
