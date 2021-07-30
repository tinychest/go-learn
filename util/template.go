package util

import (
	"strings"
	"text/template"
)

func Render(tmplStr string, params interface{}) (string, error) {
	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	return GetTemplateContent(tmpl, params)
}

func GetTemplateContent(tmpl *template.Template, params interface{}) (string, error) {
	// 输出到控制台
	// var writer io.Writer = os.Stdout
	// 输出到自定义 buffer
	var writer = &strings.Builder{}

	if err := tmpl.Execute(writer, params); err != nil {
		return "", err
	}
	return writer.String(), nil
}

func MustRender(tmplStr string, params interface{}) string {
	result, err := Render(tmplStr, params)
	if err != nil {
		panic(err)
	}
	return result
}
