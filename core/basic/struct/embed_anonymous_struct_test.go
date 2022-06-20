package _struct

import (
	"fmt"
	"testing"
)

// [内嵌匿名结构体类型]
//
// - 点
// 假如 B 声明了 F2，A 声明了 F1 和 F2 方法（F1 中调用了 F2），通过 B 调用 F1 方法和通过 A 调用 F1 方法的结果是一样的，调用的都是 A 的实现
// Go 不会像其他面向对象高级编程语言（如 Java），能够输出实际类型实现的方法
// - 点
// B 类型的变量并不能直接赋值给 A 类型的变量

type A struct {
	Name string
}

func (a A) F() {
	fmt.Println("A F", a.Name)
}

func (a *A) FP() {
	fmt.Println("A FP", a.Name)
}

type B struct {
	A
}

func TestComposite(t *testing.T) {
	var b B

	// 语法糖：A 可以直接访问内部 B 的字段 和 方法
	t.Log(b.A.Name)
	t.Log(b.Name)
	// 但是，结构体方式初始化的时候是不行的
	// var _ = B{Name:"xiaoming"}

	b.A.F()
	b.F()
	b.A.FP()
	b.FP()

	// TODO 上面按理说不是不能调用么，视频中说到该点时，涉及的 Go 版本可能比较早。下面的内容都引不出来了
	//     bp := new(B)
}
