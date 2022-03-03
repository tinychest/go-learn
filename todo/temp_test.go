package todo

import (
	"testing"
)

// package 包名如果和 Go 的关键字命名相同的话，调用时，import 正常，实际引用 Go 会自动在包名前加上 “_”

// TODO syscall.Syscall

func TestNil(t *testing.T) {
	// nil map 可以取值，不会引发 nil pointer panic
	var m map[string]int
	t.Log(m[""])

	t.Log(map[string]int(nil)[""])
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {}

func TestW(t *testing.T) {
	// live()：nil，类型：*Student（People 作为方法返回值类型可能会有包装的幌子）
	// (*Student)(nil) - 值：nil，类型：*student
	// nil - 值：nil，类型：nil
	// 所以所有接口都是不离其根的：类型 + 值
	t.Log(People((*Student)(nil)) == (*Student)(nil))
}
