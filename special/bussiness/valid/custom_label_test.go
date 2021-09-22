package valid

import (
	"reflect"
	"strings"
	"testing"
)

func TestCustomTip(t *testing.T) {
	// 自定义 tag 的值，用于自定义错误提示（字段提示首字母小写）
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 1、validate 拓展的 label
		if l := field.Tag.Get("l"); len(l) != 0 && l != "-" {
			return l
		}
		// 2、beego gin 都支持的表单参数反序列化标签
		if f := field.Tag.Get("f"); len(f) != 0 && f != "-" {
			return f
		}
		// 3、json
		if j := field.Tag.Get("json"); len(j) != 0 && j != "-" {
			return j
		}
		// 4、结构体字段名（首字母小写）
		b := strings.Builder{}
		b.Grow(len(field.Name))
		b.WriteString(strings.ToLower(string(field.Name[0])))
		b.WriteString(field.Name[1:])
		return b.String()
	})
}
