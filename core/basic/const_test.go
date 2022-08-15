package basic

import (
	"testing"
)

// 虽然基础类型的数组被归为值类型，但是并不能声明成常量

// 在同一个 const group 中，如果常量定义与前一行的定义一致，则可以省略类型和值。编译时，会按照前一行的定义自动补全
const (
	a, b = "golang", 100
	d, e
	f bool = true
	g      // goland 红线提示，实际上编译通过
)

// iota 关键字，对一组整型常量从 iota 赋值的数所在组的下标开始，默认进行每次自增 1 的值序列递增
// iota 就是 go 编译器的一个把戏，编译器会在编译的时候，根据语法规则为使用了 iota 关键字的常量组设置成具体的值
// iota 还可以组合成一个算术表达式，同时，这个表达式会应用在后续的每个数上
const (
	x = iota * 2
	y
	z
)

func TestConst(t *testing.T) {
	t.Log(a, b, d, e, f, g)
	t.Log(x, y, z)
}
