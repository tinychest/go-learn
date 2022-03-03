package basic

import (
	"testing"
)

type Int = int // 类型别名（type alias declaration）、非定义类型（undefined type）
type MyInt int // 定义类型（defined type）

// 内置类型别名（builtin.go）
// type byte = uint8
// type rune = int32

// 显示类型转换
// string ←→ []byte
// string ←→ []rune
// int ←→ float64
// int32 ←→ byte

// 列举一些可能会认为理所当然的 Go 会支持的隐式转换（实际不能，需要显示转换）
// 参考：https://gfw.go101.org/article/value-conversions-assignments-and-comparisons.html
func TestTypeCast(t *testing.T) {
	var i int
	var I Int
	var myI MyInt

	// 样例 1：定义类型 和 它的底层类型（这里和 Go白皮书 底层类型相关的类型转换规则 的说法冲突）
	// i, myI = myI, i // gg

	// 样例 2：相同底层类型的定义类型
	// type MyInt2 int
	// var _ MyInt2 = myI // gg

	// 样例 3：底层类型和它的定义类型之间可以，但是底层类型相同的两个定义类型之间不行
	// type IntSlice1 []int
	// type IntSlice2 []int
	// var is []int
	// var is1 IntSlice1
	// var is2 IntSlice2
	// is, is = is1, is2
	// is1, is2 = is, is
	// is1 = is2 // gg

	// 样例 4：https://gfw.go101.org/article/interface.html 中 "一个[]T类型的值不能直接被转换为类型[]I，即使类型T实现了接口类型I"
	// type person struct{}
	// type animal interface{}
	// var _ animal = person{}
	// var _ []animal = []person{} // gg

	// 样例 5：指针
	// type IntPtr *int
	// type MyIntPtr *MyInt
	// var pi = new(int)
	// var _ IntPtr = pi
	// // var _ *MyInt = pi // gg
	// // MyIntPtr 和 *int 类型没有任何关系，不能显示也不能隐式；但是可以间接
	// var _ MyIntPtr = (*MyInt)(pi)  // 间接隐式转换没问题
	// var _ = MyIntPtr((*MyInt)(pi)) // 间接显式转换没问题

	// 样例 6：通道
	// type C chan string
	// var c1 chan string
	// var c2 C
	// c1, c2 = c2, c1
	// type C2 chan <- string
	// type C3 <-chan string
	// var _ C2 = c1
	// var _ C3 = c1
	// var _ C2 = c2 // gg
	// var _ C3 = c2 // gg

	t.Log(i, I, myI)
}
