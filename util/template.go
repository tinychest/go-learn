package util

import (
	"strings"
	"text/template"
)

func MustRender(tplStr string, injectFuncs ...map[string]interface{}) *template.Template {
	result, err := Render(tplStr, injectFuncs...)
	if err != nil {
		panic(err)
	}
	return result
}

func MustRenderString(tplStr string, params interface{}, injectFuncs ...map[string]interface{}) string {
	result, err := RenderString(tplStr, params, injectFuncs...)
	if err != nil {
		panic(err)
	}
	return result
}

func Render(tplStr string, injectFuncs ...map[string]interface{}) (*template.Template, error) {
	tpl := template.New("test")
	if len(injectFuncs) != 0 {
		tpl.Funcs(injectFuncs[0])
	}
	return tpl.Parse(tplStr)
}

func RenderString(tplStr string, params interface{}, injectFuncs ...map[string]interface{}) (string, error) {
	tpl, err := Render(tplStr, injectFuncs...)
	if err != nil {
		return "", err
	}
	return AsString(tpl, params)
}

func AsString(tpl *template.Template, params interface{}) (string, error) {
	// 输出到控制台
	// var writer io.Writer = os.Stdout
	// 输出到自定义 buffer
	var s = &strings.Builder{}

	if err := tpl.Execute(s, params); err != nil {
		return "", err
	}
	return s.String(), nil
}
