package _struct

import "testing"

// 可以通过在 结构体 中添加一个 _ struct{} 来避免结构体的纯值实例化方式（指定字段名而是按照顺序的方式为字段赋值）
func Test(t *testing.T) {
	type S struct {
		_    struct{}
		name string
	}

	// 编译不通过
	// var s = S{"name"}

	var s = S{name: "name"}
	t.Log(s)
}
