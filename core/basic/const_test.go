package basic

import "testing"

// 在同一个 const group 中，如果常量定义与前一行的定义一致，则可以省略类型和值。编译时，会按照前一行的定义自动补全
func TestConstValue(t *testing.T) {
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g      // goland 提示编译不通过，实际上编译通过
	)
	println(d, e, g)
}
