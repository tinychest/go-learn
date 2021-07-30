package template

import (
	"fmt"
	"go-learn/util"
	"strings"
	"testing"
	"text/template"
)

const tF = `{{join . ", "}}`

func TestFuncs(t *testing.T) {
	data := []string{"A", "B", "C"}
	funcs := template.FuncMap{"join": strings.Join}

	tmpl, err := template.New("test").Funcs(funcs).Parse(tF)
	if err != nil {
		t.Fatal(err)
	}
	result, err := util.GetTemplateContent(tmpl, data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}
