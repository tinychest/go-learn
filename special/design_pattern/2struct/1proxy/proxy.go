package _proxy

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"text/template"
)

// 生成代理类的文件模板
const proxyTpl = `
package {{.Package}}

type {{ .ProxyStructName }}Proxy struct {
	child *{{ .ProxyStructName }}
}

func New{{ .ProxyStructName }}Proxy(child *{{ .ProxyStructName }}) *{{ .ProxyStructName }}Proxy {
	return &{{ .ProxyStructName }}Proxy{child: child}
}

{{ range .Methods }}
func (p *{{$.ProxyStructName}}Proxy) {{ .Name }} ({{ .Params }}) ({{ .Results }}) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	{{ .ResultNames }} = p.child.{{ .Name }}({{ .ParamNames }})

	// after 这里可能也有一些监控统计的逻辑
	fmt.Printf("user login cost time: %s", time.Now().Sub(start))

	return {{ .ResultNames }}
}
{{ end }} 
`

type proxyData struct {
	// 包名
	Package string
	// 需要代理的类名
	ProxyStructName string
	// 需要代理的方法
	Methods []*proxyMethod
}

// proxyMethod 代理的方法
type proxyMethod struct {
	// 方法名
	Name string
	// 参数，含参数类型
	Params string
	// 参数名
	ParamNames string
	// 返回值
	Results string
	// 返回值名
	ResultNames string
}

// ast：Package ast declares the types used to represent syntax trees for Go packages
// 扫码指定的静态被代理结构体，按照 proxyTpl 模板，生成得到代理类的源码
func generate(file string) (string, error) {
	var (
		f *ast.File

		tpl *template.Template
		src []byte
		err error
	)

	fset := token.NewFileSet()
	if f, err = parser.ParseFile(fset, file, nil, parser.ParseComments); err != nil {
		return "", err
	}

	// 获取代理需要的数据
	data := proxyData{
		Package: f.Name.Name,
	}

	// 构建 注释 和 node 的关系
	for node, group := range ast.NewCommentMap(fset, f, f.Comments) {
		var name string

		if name = getProxyInterfaceName(group); len(name) == 0 {
			continue
		}

		// 获取代理的类名
		data.ProxyStructName = node.(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name

		// 从文件中查找接口
		obj := f.Scope.Lookup(name)

		// 类型转换，注意：这里没有对断言进行判断，可能会导致 panic
		t := obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)

		for _, field := range t.Methods.List {
			fc := field.Type.(*ast.FuncType)

			// 代理的方法
			method := &proxyMethod{
				Name: field.Names[0].Name,
			}

			// 获取方法的参数和返回值
			method.Params, method.ParamNames = getParamsOrResults(fc.Params)
			method.Results, method.ResultNames = getParamsOrResults(fc.Results)

			data.Methods = append(data.Methods, method)
		}
	}

	// 生成文件
	if tpl, err = template.New("").Parse(proxyTpl); err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err = tpl.Execute(buf, data); err != nil {
		return "", nil
	}

	// 使用 go fmt 对生成的代码进行格式化
	if src, err = format.Source(buf.Bytes()); err != nil {
		return "", err
	}

	return string(src), nil
}

// getParamsOrResults 获取参数或是返回值
// 返回带类型的参数，以及不带类型的参数，以逗号间隔
func getParamsOrResults(fields *ast.FieldList) (string, string) {
	var (
		params     []string
		paramNames []string
	)
	for i, param := range fields.List {
		// 循环获取所有参数
		var names []string
		for _, name := range param.Names {
			names = append(names, name.Name)
		}

		if len(names) == 0 {
			names = append(names, fmt.Sprintf("r%d", i))
		}

		paramNames = append(paramNames, names...)

		// 参数名加参数类型组成完整的参数
		param := fmt.Sprintf("%s %s", strings.Join(names, ","), param.Type.(*ast.Ident).Name)
		params = append(params, strings.TrimSpace(param))
	}
	return strings.Join(params, ","), strings.Join(paramNames, ",")
}

// 遍历源码文件中所有的注释，取出格式如【@proxy 接口名】中的接口名（目前只找一个）
func getProxyInterfaceName(groups []*ast.CommentGroup) string {
	// 注释块
	for _, commentGroup := range groups {
		// 注释块中的每一条
		for _, comment := range commentGroup.List {
			if strings.Contains(comment.Text, "@proxy") {
				interfaceName := strings.TrimLeft(comment.Text, "// @proxy ")
				return strings.TrimSpace(interfaceName)
			}
		}
	}
	return ""
}
