package util

import (
	"strings"
	"text/template"
)

func MustRender(tmplStr string, injectFuncs ...map[string]interface{}) *template.Template {
	result, err := Render(tmplStr, injectFuncs...)
	if err != nil {
		panic(err)
	}
	return result
}

func MustRenderString(tmplStr string, params interface{}, injectFuncs ...map[string]interface{}) string {
	result, err := RenderString(tmplStr, params, injectFuncs...)
	if err != nil {
		panic(err)
	}
	return result
}

func Render(tmplStr string, injectFuncs ...map[string]interface{}) (*template.Template, error) {
	tpl := template.New("test")
	if len(injectFuncs) != 0 {
		tpl.Funcs(injectFuncs[0])
	}
	return tpl.Parse(tmplStr)
}

func RenderString(tmplStr string, params interface{}, injectFuncs ...map[string]interface{}) (string, error) {
	tmpl, err := Render(tmplStr, injectFuncs...)
	if err != nil {
		return "", err
	}
	return AsString(tmpl, params)
}

func AsString(tmpl *template.Template, params interface{}) (string, error) {
	// 输出到控制台
	// var writer io.Writer = os.Stdout
	// 输出到自定义 buffer
	var writer = &strings.Builder{}

	if err := tmpl.Execute(writer, params); err != nil {
		return "", err
	}
	return writer.String(), nil
}
