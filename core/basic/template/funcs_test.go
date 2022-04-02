package template

import (
	"go-learn/util"
	"strings"
	"testing"
	"text/template"
)

const tF = `{{join . ", "}}`

func TestFuncs(t *testing.T) {
	data := []string{"A", "B", "C"}
	funcs := template.FuncMap{"join": strings.Join}

	tpl, err := template.New("test").Funcs(funcs).Parse(tF)
	if err != nil {
		t.Fatal(err)
	}
	res, err := util.AsString(tpl, data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
